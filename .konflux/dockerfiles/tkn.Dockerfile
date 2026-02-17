ARG GO_BUILDER=registry.access.redhat.com/ubi9/go-toolset:9.7-1770654497@sha256:82b82ecf4aedf67c4369849047c2680dba755fe57547bbb05eca211b22038e29
ARG RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal:latest@sha256:759f5f42d9d6ce2a705e290b7fc549e2d2cd39312c4fa345f93c02e4abb8da95

FROM $GO_BUILDER AS builder

ARG REMOTE_SOURCE=/go/src/github.com/tektoncd/cli

ARG TKN_VERSION=0.41.0

WORKDIR $REMOTE_SOURCE

COPY upstream .
COPY .konflux/patches patches/
RUN set -e; for f in patches/*.patch; do echo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done

COPY head HEAD
ENV GODEBUG="http2server=0"
ENV GOEXPERIMENT=strictfipsruntime
RUN go build -ldflags="-X 'knative.dev/pkg/changeset.rev=$(cat HEAD)'" -mod=vendor -tags disable_gcp,strictfipsruntime -v \
       -ldflags "-X github.com/tektoncd/cli/pkg/cmd/version.clientVersion=${TKN_VERSION}" \
       -o /tmp/tkn ./cmd/tkn

# Build tkn-pac from sources
COPY sources/pac $REMOTE_SOURCE/pac
RUN cd $REMOTE_SOURCE/pac && \
    go build -tags strictfipsruntime -mod=vendor -o /tmp/tkn-pac ./cmd/tkn-pac

FROM $RUNTIME

ARG VERSION=tkn-next
COPY --from=builder /tmp/tkn /usr/bin
COPY --from=builder /tmp/tkn-pac /usr/bin
LABEL \
      com.redhat.component="openshift-pipelines-cli-tkn-rhel9-container" \
      cpe="cpe:/a:redhat:openshift_pipelines:1.22::el9" \
      description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
      io.k8s.description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
      io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
      io.openshift.tags="tekton,openshift,tektoncd-cli,tkn" \
      maintainer="pipelines-extcomm@redhat.com" \
      name="openshift-pipelines/pipelines-cli-tkn-rhel9" \
      summary="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
      version="v1.22.0"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532
