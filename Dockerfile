# SPDX-FileCopyrightText: 2022-present Intel Corporation
# SPDX-FileCopyrightText: 2019-present Open Networking Foundation <info@opennetworking.org>
#
# SPDX-License-Identifier: Apache-2.0

FROM onosproject/golang-build:v1.0 as build

ARG org_label_schema_version=unknown
ARG org_label_schema_vcs_url=unknown
ARG org_label_schema_vcs_ref=unknown
ARG org_label_schema_build_date=unknown
ARG org_opencord_vcs_commit_date=unknown
ARG org_opencord_vcs_dirty=unknown

ENV ADAPTER_ROOT=$GOPATH/src/github.com/onosproject/reference-adapter
ENV CGO_ENABLED=0

RUN mkdir -p $ADAPTER_ROOT/

COPY . $ADAPTER_ROOT/

RUN cat $ADAPTER_ROOT/go.mod

RUN cd $ADAPTER_ROOT && GO111MODULE=on go build -o /go/bin/reference-adapter \
        -ldflags \
        "-X github.com/onosproject/config-adapter/internal/pkg/version.Version=$org_label_schema_version \
         -X github.com/onosproject/config-adapter/internal/pkg/version.GitCommit=$org_label_schema_vcs_ref  \
         -X github.com/onosproject/config-adapter/internal/pkg/version.GitDirty=$org_opencord_vcs_dirty \
         -X github.com/onosproject/config-adapter/internal/pkg/version.GoVersion=$(go version 2>&1 | sed -E  's/.*go([0-9]+\.[0-9]+\.[0-9]+).*/\1/g') \
         -X github.com/onosproject/config-adapter/internal/pkg/version.Os=$(go env GOHOSTOS) \
         -X github.com/onosproject/config-adapter/internal/pkg/version.Arch=$(go env GOHOSTARCH) \
         -X github.com/onosproject/config-adapter/internal/pkg/version.BuildTime=$org_label_schema_build_date" \
         ./cmd/reference-adapter

FROM alpine:3.11
RUN apk add bash openssl curl libc6-compat

ENV HOME=/home/reference-adapter

RUN mkdir $HOME
WORKDIR $HOME

COPY --from=build /go/bin/reference-adapter /usr/local/bin/
