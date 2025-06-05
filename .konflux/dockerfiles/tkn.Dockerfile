ARG GO_BUILDER=brew.registry.redhat.io/rh-osbs/openshift-golang-builder:v1.22
ARG RUNTIME=registry.redhat.io/ubi8/ubi:latest@sha256:0c1757c4526cfd7fdfedc54fadf4940e7f453201de65c0fefd454f3dde117273
ARG PAC_BUILDER=registry.redhat.io/openshift-pipelines/pipelines-pipelines-as-code-cli-rhel8@sha256:b3bedee6084e9d9da549d3e637d2bf176aa365b5818e6b6377525577edd710dc

FROM $GO_BUILDER AS builder

ARG REMOTE_SOURCE=/go/src/github.com/tektoncd/cli

ARG TKN_VERSION=0.35.3

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

ARG VERSION=tkn-1.14.6
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
      io.openshift.tags="pipelines,tekton,openshift"

RUN groupadd -r -g 65532 nonroot && \
    useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532
