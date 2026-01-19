# Rebuild trigger: 1.15.4 release 2026-01-19
ARG GO_BUILDER=brew.registry.redhat.io/rh-osbs/openshift-golang-builder:v1.23
ARG RUNTIME=registry.redhat.io/ubi8/ubi:latest@sha256:8bd1b6306f8164de7fb0974031a0f903bd3ab3e6bcab835854d3d9a1a74ea5db
ARG PAC_BUILDER=quay.io/openshift-pipeline/pipelines-pipelines-as-code-cli-rhel8@sha256:9e48b69be27efc7ee21f467e68f200cb3ef103dc671e4515b06f8ac41e6c4d2e

FROM $GO_BUILDER AS builder

ARG REMOTE_SOURCE=/go/src/github.com/tektoncd/cli

ARG TKN_VERSION=0.37.2

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

ARG VERSION=tkn-1.15.3
COPY --from=builder /tmp/tkn /usr/bin
COPY --from=pacbuilder /usr/bin/tkn-pac /usr/bin
LABEL \
      com.redhat.component="openshift-pipelines-cli-tkn-container" \
      name="openshift-pipelines/pipelines-cli-tkn-rhel8" \
      version=$VERSION \
      summary="Red Hat OpenShift pipelines tkn CLI" \
      description="CLI client 'tkn' for managing openshift pipelines" \
      io.k8s.display-name="Red Hat OpenShift Pipelines tkn CLI" \
      maintainer="pipelines-extcomm@redhat.com" \
      io.k8s.description="Red Hat OpenShift Pipelines tkn CLI" \
      io.openshift.tags="pipelines,tekton,openshift" \
      cpe="cpe:/a:redhat:openshift_pipelines:1.15::el8"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot

USER 65532
