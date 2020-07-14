#!/bin/bash

set -o errexit
set -o pipefail
set -x

: "${SRC_DIR:=$(go env GOPATH)/src/github.com/kubermatic/kubermatic/api}"
: "${KUBERMATIC_WORKERNAME:=${KUBERMATIC_WORKERNAME:-$(uname -n)}}"
: "${INSTALLER_DIR:="$(go env GOPATH)/src/gitlab.syseleven.de/kubernetes/kubermatic-installer"}"
: "${KUBERMATIC_ENV:=dev}"
: "${KUBERMATIC_CLUSTER:=dbl1}"
: "${RESOURCES_DIR:=${INSTALLER_DIR}/environments/${KUBERMATIC_ENV}/clusters/${KUBERMATIC_CLUSTER}/kubermatic/versions}"
: "${CONFIG_DIR:=${INSTALLER_DIR}/environments/${KUBERMATIC_ENV}/kubermatic}"
KUBERMATIC_ENV=${KUBERMATIC_ENV} KUBERMATIC_CLUSTER=${KUBERMATIC_CLUSTER} make -C ${INSTALLER_DIR}/kubermatic multiseed_kubeconfig
: "${DEBUG:="false"}"

if [[ "${DEBUG}" == "true" ]]; then
    GOTOOLFLAGS="-v -gcflags='all=-N -l'" make -C ${SRC_DIR} master-controller-manager
else
    make -C ${SRC_DIR} master-controller-manager
fi

cd ${SRC_DIR}
if [[ "${DEBUG}" == "true" ]]; then
    dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./_build/master-controller-manager -- \
      -kubeconfig=${CONFIG_DIR}/kubeconfig \
      -internal-address=127.0.0.1:8086 \
      -worker-count=1 \
      -logtostderr=1 \
      -v=8 "$@" &
else
    ./_build/master-controller-manager \
      -kubeconfig=${CONFIG_DIR}/kubeconfig \
      -internal-address=127.0.0.1:8086 \
      -logtostderr=1 \
      -v=6 "$@" &
fi
