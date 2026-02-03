ARG GO_BUILDER=brew.registry.redhat.io/rh-osbs/openshift-golang-builder:v1.24 
ARG RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal:latest@sha256:6fc28bcb6776e387d7a35a2056d9d2b985dc4e26031e98a2bd35a7137cd6fd71
ARG PAC_BUILDER=quay.io/redhat-user-workloads/tekton-ecosystem-tenant/pipelines-pipelines-as-code-cli-rhel9@sha256:70d37fc28494e1ae13af93119333d02b6b62f6d0474653e7b7248ed5b5883cc2

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

FROM $PAC_BUILDER AS pacbuilder

FROM $RUNTIME

ARG VERSION=tkn-next
COPY --from=builder /tmp/tkn /usr/bin
COPY --from=pacbuilder /usr/bin/tkn-pac /usr/bin
LABEL \
      com.redhat.component="openshift-pipelines-cli-tkn-container" \
      name="openshift-pipelines/pipelines-cli-tkn-rhel9" \
      version=$VERSION \
      summary="Red Hat OpenShift pipelines tkn CLI" \
      description="CLI client 'tkn' for managing openshift pipelines" \
      io.k8s.display-name="Red Hat OpenShift Pipelines tkn CLI" \
      maintainer="pipelines-extcomm@redhat.com" \
      io.k8s.description="Red Hat OpenShift Pipelines tkn CLI" \
      io.openshift.tags="pipelines,tekton,openshift"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532
