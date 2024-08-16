#!/usr/bin/env zsh
BASE_DIR=".scripts"
BASE_PATH="${HOME}/${BASE_DIR}"
CONF_FILE="pkg.sh"

# Load the configuration into the installation.
source "${BASE_PATH}/${CONF_FILE}"

# Build the tools.
go build -o "${BASE_PATH}/bin/versions/${VERSION}/tools" "${BASE_PATH}/main.go"

# Make the symbolic link to the current installation version.
ln -s "${BASE_PATH}/bin/versions/${VERSION}/tools" "${BASE_PATH}/bin/local/tools"

echo "Done"
