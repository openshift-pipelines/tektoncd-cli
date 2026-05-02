# Rebuild trigger: 1.15.4 release 2026-02-27
ARG GO_BUILDER=registry.access.redhat.com/ubi9/go-toolset:1.25.9-1777537863
ARG RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal:latest
ARG PAC_BUILDER=registry.redhat.io/openshift-pipelines/pipelines-pipelines-as-code-cli-rhel8@sha256:c3c8fcadb15bed1134934a603264cdab044a356102c5330c0f2923c5588c71ef

FROM $GO_BUILDER AS builder

ARG REMOTE_SOURCE=/go/src/github.com/tektoncd/cli

ARG TKN_VERSION=0.37.3

WORKDIR $REMOTE_SOURCE

COPY upstream .
COPY .konflux/patches patches/
RUN set -e; for f in patches/*.patch; do echo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done

COPY head HEAD
ENV GODEBUG="http2server=0"
RUN go build -ldflags="-X 'knative.dev/pkg/changeset.rev=$(cat HEAD)'" -mod=vendor -tags disable_gcp -v \
       -ldflags "-X github.com/tektoncd/cli/pkg/cmd/version.clientVersion=${TKN_VERSION}" \
       -o /tmp/tkn ./cmd/tkn

FROM $PAC_BUILDER AS pacbuilder

FROM $RUNTIME

ARG VERSION=1.15
COPY --from=builder /tmp/tkn /usr/bin
COPY --from=pacbuilder /usr/bin/tkn-pac /usr/bin
LABEL \
    com.redhat.component="openshift-pipelines-cli-tkn-rhel9-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.15::el9" \
    description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    io.openshift.tags="tekton,openshift,tektoncd-cli,tkn" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-cli-tkn-rhel9" \
    summary="Red Hat OpenShift Pipelines tektoncd-cli tkn" \
    version="v1.15.5"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot

USER 65532
# trigger rebuild 2026-02-14
