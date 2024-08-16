function nvm-version() {
  echo "+ nvm    | $(nvm --version)";
  echo "|-> node | $(node --version)";
  echo "|-> npm  | $(npm --version)";
  echo "|-> yarn | $(yarn --version)";
  echo "\`-> pnpm | $(pnpm --version)";
}

function nvm-install-lts() {
  echo "nvm $(nvm --version)"
  echo "\n[BEFORE]"
  nvm-version;
  printf "\n"

  nvm install --lts;
  nvm use --lts --default;

  npm i -g npm;
  npm i -g yarn;
  npm i -g pnpm;

  echo "\n[AFTER]"
  nvm-version;
}
