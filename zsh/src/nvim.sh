function nvim-reinstall() {
  # Clean the current config and caches out.
  rm -rfv ~/.config/nvim ~/.local/share/nvim ~/.cache/nvim;

  # Download a new copy of the NVIM config.
  gh repo clone jgttech/nvim ${HOME}/.config/nvim;

  npm i -g typescript typescript-language-server;

  nvim;
}

function nvim-sync() {
  cwd=$(pwd)

  cd ${HOME}/.config/nvim;
  git pull

  cd $cwd
  unset cwd
}

function nvim-edit() {
  nvim ${HOME}/.config/nvim
}

function nvim-scripts() {
  nvim ${HOME}/.scripts
}

function nvim-update() {
  timestamp=$(date +%F\ %T)
  cwd=$(pwd)

  cd ${HOME}/.config/nvim
  git add .;
  git commit -m "${timestamp} - UPDATES";
  git push;

  cd $cwd
  unset cwd
}

