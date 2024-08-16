#!/usr/bin/env zsh
BASE_DIR=".tools"
BASE_PATH="${HOME}/${BASE_DIR}"
CONF_PATH="${BASE_PATH}/pkg.sh"

# Load the configuration into the installation.
source ${CONF_PATH}

# Clone the repo using the GitHub CLI.
gh repo clone jgttech/tools ${BASE_PATH}

# Build the tools.
go build -o "${BASE_PATH}/bin/versions/${VERSION}/tools" "${BASE_PATH}/main.go"

# Make the symbolic link to the current installation version.
ln -s "${BASE_PATH}/bin/versions/${VERSION}/tools" "${BASE_PATH}/bin/local/tools"

echo "Done"
