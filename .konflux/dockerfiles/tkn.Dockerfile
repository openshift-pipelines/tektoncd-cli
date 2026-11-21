ARG GO_BUILDER=registry.redhat.io/ubi9/go-toolset:latest@sha256:51eef2836e5223c41577757c30544fbfb3f0710f3436b70b82d2195e35fcb743
ARG UBI_RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal:latest@sha256:08d4c70f1baf6537a0728656cf7942d561e74a35916509d205b6519861266352

FROM $GO_BUILDER AS builder

ARG REMOTE_SOURCE=/go/src/github.com/tektoncd/cli


WORKDIR $REMOTE_SOURCE

COPY upstream .
COPY .konflux/patches patches/
RUN set -e; for f in patches/*.patch; do echo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done

COPY head HEAD
ENV GODEBUG="http2server=0"
ENV GOEXPERIMENT=strictfipsruntime
RUN TKN_VERSION=$(cat VERSION);\
    go build -ldflags="-X 'knative.dev/pkg/changeset.rev=$(cat HEAD)'" -mod=vendor -tags disable_gcp,strictfipsruntime -v \
      -ldflags "-X github.com/tektoncd/cli/pkg/cmd/version.clientVersion=${TKN_VERSION}" \
      -o /tmp/tkn ./cmd/tkn

# Build tkn-pac from sources
COPY sources/pac $REMOTE_SOURCE/pac
RUN cd $REMOTE_SOURCE/pac && \
    go build -tags strictfipsruntime -mod=vendor -o /tmp/tkn-pac ./cmd/tkn-pac

FROM $UBI_RUNTIME

ARG VERSION=1.23
COPY --from=builder /tmp/tkn /usr/bin
COPY --from=builder /tmp/tkn-pac /usr/bin
LABEL \
    com.redhat.component="openshift-pipelines-cli-tkn-rhel9-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.23::el9" \
    description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.openshift.tags="tekton,openshift,tektoncd-cli,tkn" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-cli-tkn-rhel9" \
    summary="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    version="v1.23.5"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532
