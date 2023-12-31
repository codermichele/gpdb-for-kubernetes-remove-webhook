# syntax = docker/dockerfile:experimental
FROM golang:1.19 as build-in-docker

COPY greenplum-operator/ /greenplum-for-kubernetes/greenplum-operator/
COPY pkg/ /greenplum-for-kubernetes/pkg/
COPY go.mod go.sum /greenplum-for-kubernetes/
COPY greenplum-operator/vendor /greenplum-for-kubernetes/vendor
ENV TZ=Asia/Shanghai LANG=en_US.UTF-8
#WORKDIR /greenplum-for-kubernetes
#RUN go mod vendor \
#&& find / -name vendor

WORKDIR /greenplum-for-kubernetes/greenplum-operator
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    cd cmd/greenplumOperator && go build

# build greenplum-operator image from here
FROM cloudfoundry/run:base

COPY --from=build-in-docker \
    /greenplum-for-kubernetes/greenplum-operator/cmd/greenplumOperator/greenplumOperator \
    /usr/local/bin/greenplum-operator

#RUN set -x && \
#    apt-get update && \
#    apt-get upgrade -y && \
#    apt-get clean && \
#    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
COPY --from=build-in-docker \
     /greenplum-for-kubernetes/greenplum-operator/greenplum_operator_apt_install_order.txt \
     /greenplum-for-kubernetes/greenplum-operator/greenplum-operator-debs \
     /greenplum-for-kubernetes/greenplum-operator/sysconf.bash \
     /greenplum-for-kubernetes/greenplum-operator/

RUN set -x \
&& chmod +x /greenplum-for-kubernetes/greenplum-operator/sysconf.bash \
&& /greenplum-for-kubernetes/greenplum-operator/sysconf.bash \
&& rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

ARG DATE="10 July 2023"
ARG LICENSE="https://www.apache.org/licenses/LICENSE-2.0"
ARG GREENPLUM_VERSION=6.24.6

ENV GREENPLUM_VERSION=${GREENPLUM_VERSION}

LABEL name="Greenplum Operator" \
    vendor="jike21" \
    version="${VERSION}" \
    build_date="${DATE}" \
    license="${LICENSE}" \
    greenplum_version="${GREENPLUM_VERSION}"
