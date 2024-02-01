#!/bin/sh

OUT_DIR=..
PACK_NAME='models'
SUFFIX_NAME='do'
READ_ONLY='created_time,updated_time'
DSN_URL='mysql://root:123456@192.168.1.16:3306/coin-ant?charset=utf8'
SPEC_TYPES=''
TINYINT_TO_BOOL='deleted,is_ok'
IMPORT_MODELS='coin-ant/models'
db2go --url "${DSN_URL}" --out "${OUT_DIR}" --enable-decimal --spec-type "${SPEC_TYPES}" \
      --suffix "${SUFFIX_NAME}" --package "${PACK_NAME}" --readonly "${READ_ONLY}"  --tinyint-as-bool "${TINYINT_TO_BOOL}" \
      --dao dao --import-models "${IMPORT_MODELS}"

gofmt -w "${OUT_DIR}"/"${PACK_NAME}"
