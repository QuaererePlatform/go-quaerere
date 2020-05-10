#!/usr/bin/env sh

set -e

if [ "${COLUMBIA_DISABLE_TLS}" != "true" ]; then
  if [ ! -z ${TLS_CERT+x} ] && [ ! -z ${COLUMBIA_TLS_CERT_PATH+x} ]; then
    mkdir -p "$(dirname ${COLUMBIA_TLS_CERT_PATH})"
    echo "${TLS_CERT}" > ${COLUMBIA_TLS_CERT_PATH}
  fi

  if [ ! -z ${TLS_KEY+x} ] && [ ! -z ${COLUMBIA_TLS_KEY_PATH+x} ]; then
    mkdir -p "$(dirname ${COLUMBIA_TLS_KEY_PATH})"
    echo "${TLS_KEY}" > ${COLUMBIA_TLS_KEY_PATH}
  fi
else
  echo "[INFO] skipping TLS configuration"
fi

exec /opt/columbia/columbia
