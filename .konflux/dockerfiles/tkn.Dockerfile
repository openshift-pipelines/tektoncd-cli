ARG GO_BUILDER=registry.redhat.io/ubi9/go-toolset:latest@sha256:a2ad457dc7525635613d2355037239ca9bf4374ba3b59961d1cdaa4d5f1af7d8
ARG UBI_RUNTIME=registry.redhat.io/ubi9/ubi-minimal-pqc:latest@sha256:9aca95ae4e261eab848dc6abdfd08eb4c2453944232021245b2e3108eb25ae18

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

FROM $UBI_RUNTIME

ARG VERSION=next
COPY --from=builder /tmp/tkn /usr/bin
COPY --from=builder /tmp/tkn-pac /usr/bin
LABEL \
    com.redhat.component="openshift-pipelines-cli-tkn-rhel9-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:next::el9" \
    description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.openshift.tags="tekton,openshift,tektoncd-cli,tkn" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-cli-tkn-rhel9" \
    summary="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    version="next"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532
