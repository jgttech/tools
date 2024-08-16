#!/usr/bin/env zsh
BASE_DIR=".tools"
LINK_NAME="tools"
BASE_PATH="${HOME}/${BASE_DIR}"
CONF_PATH="${BASE_PATH}/pkg.sh"
LOCAL_PATH="${BASE_PATH}/bin/local"
LOCAL_LINK="${BASE_PATH}/bin/local/${LINK_NAME}"

# Clone the repo using the GitHub CLI.
gh repo clone jgttech/tools ${BASE_PATH}

# Load the configuration into the installation.
source ${CONF_PATH}

if [ ! -d ${LOCAL_PATH} ]; then
  mkdir ${LOCAL_PATH}
fi

# Build the tools.
go build -o "${BASE_PATH}/bin/versions/${VERSION}/tools" "${BASE_PATH}/main.go"

# Make the symbolic link to the current installation version.
# ln -s "${BASE_PATH}/bin/versions/${VERSION}/tools" "${BASE_PATH}/bin/local/tools"

echo "Done"
