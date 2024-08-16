#!/usr/bin/env zsh
BASE_DIR=".scripts"
BASE_PATH="${HOME}/${BASE_DIR}"
CONF_FILE="pkg.sh"

source "${BASE_PATH}/${CONF_FILE}"

cd ${BASE_PATH}
go build -o "${BASE_PATH}/bin/${VERSION}/tools" "${BASE_PATH}/main.go"

echo "Done"
