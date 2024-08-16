#!/usr/bin/env zsh
BASE_DIR=".tools"
LINK_NAME="tools"
BASE_PATH="${HOME}/${BASE_DIR}"
CONF_PATH="${BASE_PATH}/pkg.sh"
LOCAL_PATH="${BASE_PATH}/bin/local"
LOCAL_LINK="${LOCAL_PATH}/${LINK_NAME}"

# Clone the repo using the GitHub CLI.
gh repo clone jgttech/tools ${BASE_PATH}

# Load the configuration into the installation.
source ${CONF_PATH}

# Configuration-based envronment configuration.
BIN_PATH="${BASE_PATH}/bin/versions/${VERSION}/tools"

if [ ! -d ${LOCAL_PATH} ]; then
  mkdir ${LOCAL_PATH}
  ln -s ${BIN_PATH} ${LOCAL_LINK}
fi

# Build the tools.
go build -o ${BIN_PATH} "${BASE_PATH}/main.go"

${LOCAL_LINK}

echo "Done"
