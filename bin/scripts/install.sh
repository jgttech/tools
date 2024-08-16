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

if [ -d ${BASE_PATH} ]; then
  echo "Looks like this already exists: '${BASE_PATH}'\n"
  echo "[NOTICE]"
  echo "You can remove this directory to re-install the tools.\n"
  echo "rm -rfv ${BASE_PATH}\n"
  echo "Then run the install command again."
  exit
fi

# Clone the repo using the GitHub CLI.
gh repo clone jgttech/tools ${BASE_PATH}

# Load the configuration into the installation.
source ${CONF_PATH}

# Configuration-based envronment configuration.
BIN_PATH="${BASE_PATH}/bin/versions/${VERSION}/tools"
SEARCH_CRITERIA="\${HOME}/${BASE_DIR}/bin/local"


# PROFILE_FILE=".zshrc"
# BASE_DIR=".tools"
# LINK_NAME="tools"
# BASE_PATH="${HOME}/${BASE_DIR}"
# CONF_PATH="${BASE_PATH}/pkg.sh"
# LOCAL_PATH="${BASE_PATH}/bin/local"
# LOCAL_LINK="${LOCAL_PATH}/${LINK_NAME}"
# PROFILE_PATH="${HOME}/${PROFILE_FILE}"
#
# # Clone the repo using the GitHub CLI.
# gh repo clone jgttech/tools ${BASE_PATH}
#
# # Load the configuration into the installation.
# source ${CONF_PATH}
#
# # Configuration-based envronment configuration.
# BIN_PATH="${BASE_PATH}/bin/versions/${VERSION}/tools"
# GREP_CRITERIA="\${HOME}/${BASE_DIR}/bin/local"
#
# if [ ! -d ${LOCAL_PATH} ]; then
#   mkdir ${LOCAL_PATH}
#   ln -s ${BIN_PATH} ${LOCAL_LINK}
# fi
#
# # Download the modules
# cd ${BASE_PATH}
# go mod download
#
# # Build the tools.
# go build -o ${BIN_PATH} "${BASE_PATH}/main.go"
#
# function tools {
#   ${LOCAL_LINK} $@
# }
#
# case `grep -Fq ${GREP_CRITERIA} ${PROFILE_PATH} >/dev/null; echo $?` in
#   0)
#     # Code is found
#     echo "\n+------------------------------------------+"
#     echo "| (i) INFO                                 |"
#     echo "|                                          |"
#     echo "| Looks like you already have the tools    |"
#     echo "| linked in your \${HOME}/.zshrc config.    |"
#     echo "+------------------------------------------+\n"
#     ;;
#   1)
#     # Code if not found
#     unix_timestamp=`date +%s`
#     timestamp=`date`
#
#     link="# jgttech/tools (generated on ${timestamp})\n"
#     link="${link}if [ -d \${HOME}/${BASE_DIR} ]; then\n"
#     link="${link}  # Add the symbolic link to the version binary to the shell PATH.\n"
#     link="${link}  export PATH=\"\${HOME}/${BASE_DIR}/bin/local:\${PATH}\"\n\n"
#     link="${link}  # Used to link any shell configuration(s).\n"
#     link="${link}  source \${HOME}/${BASE_DIR}/zshrc.sh\n"
#     link="${link}fi\n"
#     link="${link}\n$(cat ${PROFILE_PATH})"
#
#     backup="${PROFILE_FILE}.${unix_timestamp}.bak"
#
#     echo "Creating profile backup here: \"\${HOME}/${backup}\""
#     cp ${PROFILE_PATH} ${backup}
#
#     echo "Updating profile to link tools"
#     echo "${m}" > ${PROFILE_PATH}
#
#     case `grep -Fq ${GREP_CRITERIA} ${PROFILE_PATH} >/dev/null; echo $?` in
#       0)
#         # Code is found
#         echo "\n+------------------------------------------+"
#         echo "| (+) SUCCESS                              |"
#         echo "|                                          |"
#         echo "| Sucessfully linked tools to your config. |"
#         echo "+------------------------------------------+\n"
#         ;;
#       1)
#         # Code is not found
#         echo "\n+------------------------------------------+"
#         echo "| (-) FAILURE                              |"
#         echo "|                                          |"
#         echo "| Failed to link tools to your config.     |"
#         echo "+------------------------------------------+\n"
#         ;;
#       *)
#         echo "\n+------------------------------------------+"
#         echo "| {!} ERROR (2)                            |"
#         echo "|                                          |"
#         echo "| Oops, looks like something went wrong.   |"
#         echo "| Failed to link and/or install the tools. |"
#         echo "+------------------------------------------+\n"
#         ;;
#     esac
#     ;;
#   *)
#     # Code if an error occurred
#     echo "\n+------------------------------------------+"
#     echo "| {!} ERROR (1)                            |"
#     echo "|                                          |"
#     echo "| Oops, looks like something went wrong.   |"
#     echo "| Failed to link and/or install the tools. |"
#     echo "+------------------------------------------+\n"
#     ;;
# esac
#
# echo "Done"
