ARG GO_BUILDER=registry.access.redhat.com/ubi10/go-toolset:latest@sha256:cb712e78c45e429142e846c1b8b91894970eab3b66fd5a95d336b65b4f7bb877
ARG UBI_RUNTIME=registry.access.redhat.com/ubi10/ubi-minimal:latest@sha256:c3fe95ec9f53a3888c47314eb4d08bde09ebde646a428137b445091afe1daf89

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

ARG VERSION=nightly
COPY --from=builder /tmp/tkn /usr/bin
COPY --from=builder /tmp/tkn-pac /usr/bin
LABEL \
    com.redhat.component="openshift-pipelines-cli-tkn-rhel10-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:nightly::el10" \
    description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.openshift.tags="tekton,openshift,tektoncd-cli,tkn" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-cli-tkn-rhel10" \
    summary="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    version="nightly"

# ubi10/ubi-minimal does not ship shadow-utils, which provides groupadd/useradd
RUN microdnf install -y shadow-utils && \
    groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532
