ARG GO_BUILDER=registry.access.redhat.com/ubi9/go-toolset:1.25.8-1776962329
ARG RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal:latest

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
    echo "Build TKN ($TKN_VER)" ;\
    go build -ldflags="-X 'knative.dev/pkg/changeset.rev=$(cat HEAD)'" -mod=vendor -tags disable_gcp,strictfipsruntime -v \
       -ldflags "-X github.com/tektoncd/cli/pkg/cmd/version.clientVersion=${TKN_VERSION}" \
       -o /tmp/tkn ./cmd/tkn

# Build tkn-pac from sources
COPY sources/pac $REMOTE_SOURCE/pac
RUN PAC_VER=$(cat pac/pkg/params/version/version.txt); \
    echo "Build TKN-PAC Version : ($PAC_VER)"; \
    cd $REMOTE_SOURCE/pac ; \
    go build -tags strictfipsruntime -mod=vendor \
        -ldflags "-X github.com/tektoncd/pipelines-as-code/pkg/params/version.Version=${PAC_VER}"\
        -o /tmp/tkn-pac ./cmd/tkn-pac

FROM $RUNTIME

COPY --from=builder /tmp/tkn /usr/bin
COPY --from=builder /tmp/tkn-pac /usr/bin
LABEL \
    com.redhat.component="openshift-pipelines-cli-tkn-rhel9-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.21::el9" \
    description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.openshift.tags="tekton,openshift,tektoncd-cli,tkn" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-cli-tkn-rhel9" \
    summary="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    version="v1.21.1"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532
