# Rebuild trigger: 1.15.4 release 2026-01-19
ARG GO_BUILDER=registry.access.redhat.com/ubi8/go-toolset:1.25.7-1771287729
ARG RUNTIME=registry.redhat.io/ubi8/ubi@sha256:87463a8cd4ea7b3e7d066f114b64a44782515031b221253d1357d27572b6d53e
ARG PAC_BUILDER=quay.io/openshift-pipeline/pipelines-pipelines-as-code-cli-rhel8@sha256:4fe9af08579f18a9389fa3d142b384eb5c88a49b702d11720868595f3be52fa5

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

ARG VERSION=tkn-1.15.4
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
# trigger rebuild 2026-02-14
