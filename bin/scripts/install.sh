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

# Setup configuration
LOCAL_PATH="${BASE_PATH}/${LOCAL_DIR}"
LOCAL_LINK="${LOCAL_PATH}/${NAME}"
BIN_PATH="${BASE_PATH}/${VERSIONS_DIR}/${VERSION}/${NAME}"
SEARCH_CRITERIA="\${HOME}/${BASE_DIR}/${LOCAL_DIR}"
PROFILE_PATH="${HOME}/${PROFILE}"

if [ ! -d ${LOCAL_PATH} ]; then
  mkdir ${LOCAL_PATH}
  ln -s ${BIN_PATH} ${LOCAL_LINK}
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
    link="${link}  export PATH=\"\${HOME}/${BASE_DIR}/bin/local:\${PATH}\"\n\n"
    link="${link}  # Used to link any shell configuration(s).\n"
    link="${link}  source \${HOME}/${BASE_DIR}/zshrc.sh\n"
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
        echo "| Sucessfully linked tools to your config."
        ;;
      1)
        # Code is not found
        echo "\n| {!} FAILURE"
        echo "|"
        echo "| Failed to link tools to your config."
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

echo "Done"
