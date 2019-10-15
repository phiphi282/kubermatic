#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail
set -x

: "${SRC_DIR:=$(go env GOPATH)/src/github.com/kubermatic/kubermatic/api}"
: "${KUBERMATIC_WORKERNAME:=${KUBERMATIC_WORKERNAME:-$(uname -n)}}"
: "${INSTALLER_DIR:="$(go env GOPATH)/src/gitlab.syseleven.de/kubernetes/kubermatic-installer"}"
: "${KUBERMATIC_ENV:=dev}"
: "${KUBERMATIC_CLUSTER:=dbl1}"
: "${RESOURCES_DIR:=${INSTALLER_DIR}/environments/${KUBERMATIC_ENV}/clusters/${KUBERMATIC_CLUSTER}/kubermatic/versions}"
: "${CONFIG_DIR:=${INSTALLER_DIR}/environments/${KUBERMATIC_ENV}/kubermatic}"
KUBERMATIC_ENV=${KUBERMATIC_ENV} KUBERMATIC_CLUSTER=${KUBERMATIC_CLUSTER} make -C ${INSTALLER_DIR}/kubermatic values.yaml
: "${EXTERNAL_URL:=dev.metakube.de}"
: "${DEBUG:="false"}"

while true; do
    if [[ "${DEBUG}" == "true" ]]; then
        GOTOOLFLAGS="-v -gcflags='all=-N -l'" make -C ${SRC_DIR} kubermatic-controller-manager
    else
        make -C ${SRC_DIR} kubermatic-controller-manager
    fi

    cd ${SRC_DIR}
    if [[ "${DEBUG}" == "true" ]]; then
        dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ./_build/kubermatic-controller-manager \
          -datacenters=${CONFIG_DIR}/datacenters.yaml \
          -datacenter-name=${KUBERMATIC_CLUSTER} \
          -kubeconfig=${CONFIG_DIR}/kubeconfig \
          -versions=${RESOURCES_DIR}/versions.yaml \
          -updates=${RESOURCES_DIR}/updates.yaml \
          -master-resources=${RESOURCES_DIR} \
          -kubernetes-addons-path=${INSTALLER_DIR}/kubermatic/cluster-addons/addons \
          -openshift-addons-path=../openshift_addons \
          -worker-name="$(tr -cd '[:alnum:]' <<< $KUBERMATIC_WORKERNAME | tr '[:upper:]' '[:lower:]')" \
          -external-url=${EXTERNAL_URL} \
          -docker-pull-config-json-file=${INSTALLER_DIR}/kubermatic/dockerconfigjson \
          -monitoring-scrape-annotation-prefix=${KUBERMATIC_ENV} \
          -logtostderr=1 \
          -v=8 $@ &

        PID=$!
    else
        ./_build/kubermatic-controller-manager \
          -datacenters=${CONFIG_DIR}/datacenters.yaml \
          -datacenter-name=${KUBERMATIC_CLUSTER} \
          -kubeconfig=${CONFIG_DIR}/kubeconfig \
          -versions=${RESOURCES_DIR}/versions.yaml \
          -updates=${RESOURCES_DIR}/updates.yaml \
          -master-resources=${RESOURCES_DIR} \
          -kubernetes-addons-path=${INSTALLER_DIR}/kubermatic/cluster-addons/addons \
          -openshift-addons-path=../openshift_addons \
          -worker-name="$(tr -cd '[:alnum:]' <<< $KUBERMATIC_WORKERNAME | tr '[:upper:]' '[:lower:]')" \
          -external-url=${EXTERNAL_URL} \
          -docker-pull-config-json-file=${INSTALLER_DIR}/kubermatic/dockerconfigjson \
          -monitoring-scrape-annotation-prefix=${KUBERMATIC_ENV} \
          -logtostderr=1 \
          -v=6 $@ &

          # TODO
          #-backup-container=../config/kubermatic/static/backup-container.yaml \
          #-cleanup-container=../config/kubermatic/static/cleanup-container.yaml \
          #-oidc-ca-file=../../secrets/seed-clusters/dev.kubermatic.io/caBundle.pem \
          #-oidc-issuer-url=$(vault kv get -field=oidc-issuer-url dev/seed-clusters/dev.kubermatic.io) \
          #-oidc-issuer-client-id=$(vault kv get -field=oidc-issuer-client-id dev/seed-clusters/dev.kubermatic.io) \
          #-oidc-issuer-client-secret=$(vault kv get -field=oidc-issuer-client-secret dev/seed-clusters/dev.kubermatic.io) \
          #-docker-pull-config-json-file=<generate the file>

        PID=$!
    fi



    if [[ -x "$(command -v inotifywait)" ]]; then
        inotifywait -r -e modify ${SRC_DIR}
    elif [[ -x "$(command -v fswatch)" ]]; then
        fswatch --one-event ${SRC_DIR}
    else
        echo "Can not watch changes because neither fswatch (MAC) nor inotifywait found"
        kill ${PID}
        exit 1
    fi


    echo "Change in kubermatic detected, recompiling and restarting"

    kill ${PID} || true
done