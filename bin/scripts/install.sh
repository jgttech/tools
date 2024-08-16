#!/usr/bin/env zsh
PROFILE_FILE=".zshrc"
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
GREP_CRITERIA="\${HOME}/${BASE_DIR}/bin/local"

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

case `grep -Fq ${GREP_CRITERIA} "${HOME}/${PROFILE_FILE}" >/dev/null; echo $?` in
  0)
    # Code is found
    echo "\n+------------------------------------------+"
    echo "| (i) INFO                                 |"
    echo "|                                          |"
    echo "| Looks like you already have the tools    |"
    echo "| linked in your \${HOME}/.zshrc config.    |"
    echo "+------------------------------------------+\n"
    ;;
  1)
    # Code if not found
    timestamp=`date +%s`

    m="# jgttech/tools\n"
    m="${m}if [ -d \${HOME}/${BASE_DIR} ]; then\n"
    m="${m}  export PATH=\"\${HOME}/${BASE_DIR}/bin/local:\${PATH}\"\n"
    m="${m}  tools sync\n"
    m="${m}fi\n"
    m="${m}\n$(cat ${HOME}/${PROFILE_FILE})"

    # echo "${m}" > "${HOME}/${PROFILE_FILE}"
    echo "${m}"

    case `grep -Fq ${GREP_CRITERIA} "${HOME}/${PROFILE_FILE}" >/dev/null; echo $?` in
      0)
        # Code is found
        echo "\n+------------------------------------------+"
        echo "| (+) SUCCESS                              |"
        echo "|                                          |"
        echo "| Sucessfully linked tools to your config. |"
        echo "+------------------------------------------+\n"
        ;;
      1)
        # Code is not found
        echo "\n+------------------------------------------+"
        echo "| (-) FAILURE                              |"
        echo "|                                          |"
        echo "| Failed to link tools to your config.     |"
        echo "+------------------------------------------+\n"
        ;;
      *)
        echo "\n+------------------------------------------+"
        echo "| {!} ERROR (2)                            |"
        echo "|                                          |"
        echo "| Oops, looks like something went wrong.   |"
        echo "| Failed to link and/or install the tools. |"
        echo "+------------------------------------------+\n"
        ;;
    esac
    ;;
  *)
    # Code if an error occurred
    echo "\n+------------------------------------------+"
    echo "| {!} ERROR (1)                            |"
    echo "|                                          |"
    echo "| Oops, looks like something went wrong.   |"
    echo "| Failed to link and/or install the tools. |"
    echo "+------------------------------------------+\n"
    ;;
esac

echo "Done"
