#!/usr/bin/env sh

set -e

if [ "${KOOTENAY_DISABLE_TLS}" != "true" ]; then
  if [ ! -z ${TLS_CERT+x} ] && [ ! -z ${KOOTENAY_TLS_CERT_PATH+x} ]; then
    mkdir -p "$(dirname ${KOOTENAY_TLS_CERT_PATH})"
    echo "${TLS_CERT}" > ${KOOTENAY_TLS_CERT_PATH}
  fi

  if [ ! -z ${TLS_KEY+x} ] && [ ! -z ${KOOTENAY_TLS_KEY_PATH+x} ]; then
    mkdir -p "$(dirname ${KOOTENAY_TLS_KEY_PATH})"
    echo "${TLS_KEY}" > ${KOOTENAY_TLS_KEY_PATH}
  fi
else
  echo "[INFO] skipping TLS configuration"
fi

exec /opt/kootenay/kootenay
