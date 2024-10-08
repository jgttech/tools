#!/usr/bin/env zsh
function trim {
    local var="$*"
    var="${var#"${var%%[![:space:]]*}"}"
    var="${var%"${var##*[![:space:]]}"}"
    printf '%s' "$var"
}

BASE_DIR=$(trim $1)
BASE_PATH="${HOME}/${BASE_DIR}"
CONF_PATH="${BASE_PATH}/pkg.sh"

# Tools installation check
if [ -d ${BASE_PATH} ]; then
  echo "\nLooks like this already exists: '${BASE_PATH}'"
  echo "You can delete this directory to re-install the tools."
  exit
fi

# Clone the repo using the GitHub CLI.
gh repo clone jgttech/tools ${BASE_PATH}

# Load the configuration into the installation.
source ${CONF_PATH}

# The output directory for all the generated code.
OUT_PATH="${BASE_PATH}/${OUT_DIR}"

# Create the output directory, if it does not exist.
if [ ! -d ${OUT_PATH} ]; then
  mkdir ${OUT_PATH}
fi

# Setup configuration
LOCAL_PATH="${OUT_PATH}/${LOCAL_DIR}"
LOCAL_LINK="${LOCAL_PATH}/${NAME}"
BIN_PATH="${OUT_PATH}/${VERSIONS_DIR}/${VERSION}/${NAME}"
SEARCH_CRITERIA="\${HOME}/${BASE_DIR}/${OUT_DIR}/${LOCAL_DIR}"
PROFILE_PATH="${HOME}/${PROFILE}"

if [ ! -d ${LOCAL_PATH} ]; then
  mkdir ${LOCAL_PATH}
  ln -s ${BIN_PATH} ${LOCAL_LINK}
fi

GO_ENV_DIR="${OUT_PATH}/env"
GO_ENV_FILE="${GO_ENV_DIR}/env.go"

if [ ! -d ${GO_ENV_DIR} ]; then
  mkdir ${GO_ENV_DIR}
  touch ${GO_ENV_FILE}

  # Generates an "env" module in Go that gets
  # compiled into the binary and used to know
  # what the installation environment was/is.
  env="package env\n\n"
  env="${env}const (\n"
  env="${env}  BASE_DIR     = \"${BASE_DIR}\"\n"
  env="${env})\n"

  echo "${env}" > ${GO_ENV_FILE}
fi

# Download the modules
cd ${BASE_PATH}
go mod download

# Build the tools.
go build -o ${BIN_PATH} "${BASE_PATH}/main.go"

function tools {
  ${LOCAL_LINK} $@
}

# Check the
case `grep -Fq ${SEARCH_CRITERIA} ${PROFILE_PATH} >/dev/null; echo $?` in
  0)
    # Code is found
    echo "\n| (i) INFO|"
    echo "|"
    echo "| Looks like you already have the tools"
    echo "| linked in your \${HOME}/${PROFILE} config."
    ;;
  1)
    # Code if not found
    unix_timestamp=`date +%s`
    timestamp=`date`

    link="# jgttech/tools (generated on ${timestamp})\n"
    link="${link}if [ -d \${HOME}/${BASE_DIR} ]; then\n"
    link="${link}  # Add the symbolic link to the version binary to the shell PATH.\n"
    link="${link}  export PATH=\"\${HOME}/${BASE_DIR}/${OUT_DIR}/${LOCAL_DIR}:\${PATH}\"\n\n"
    link="${link}  # Used to link any shell configuration(s).\n"
    link="${link}  source \${HOME}/${BASE_DIR}/bin/${SHELL_DIR}/${PROFILE}\n"
    link="${link}fi\n"
    link="${link}\n$(cat ${PROFILE_PATH})"

    backup="${HOME}/${PROFILE}.${unix_timestamp}.bak"

    echo "Creating profile backup here: \"\${HOME}/${backup}\""
    cp ${PROFILE_PATH} ${backup}

    echo "Updating profile to link tools"
    echo "${link}" > ${PROFILE_PATH}

    case `grep -Fq ${SEARCH_CRITERIA} ${PROFILE_PATH} >/dev/null; echo $?` in
      0)
        # Code is found
        echo "\n| (+) SUCCESS"
        echo "|"
        echo "| Sucessfully linked tools to your config.\n"

        echo "Closing the terminal and restarting is should"
        echo "successfully, allow you to run 'tools' and see"
        echo "the CLI help information."
        ;;
      1)
        # Code is not found
        echo "\n| {!} WARNING"
        echo "|"
        echo "| Failed to link tools to your config."
        echo "| You might already have a tools config"
        echo "| already setup."
        ;;
      *)
        # Code if an error occurred
        echo "\n| /!\\ ERROR"
        echo "|"
        echo "| Oops, looks like something went wrong."
        echo "| Failed to link and/or install the tools."
        ;;
    esac
    ;;
  *)
    # Code if an error occurred
    echo "\n| /!\\ ERROR"
    echo "|"
    echo "| Oops, looks like something went wrong."
    echo "| Failed to link and/or install the tools."
    ;;
esac

source ${HOME}/${PROFILE}
echo "\nDone"
