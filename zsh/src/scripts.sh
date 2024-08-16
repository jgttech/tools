function scripts-sync() {
  cwd=$(pwd)

  cd ${HOME}/.scripts;
  git pull

  cd $cwd
  unset cwd
}

function scripts-update() {
  timestamp=$(date +%F\ %T)
  cwd=$(pwd)

  cd ${HOME}/.scripts
  git add .;
  git commit -m "${timestamp} - UPDATES";
  git push;

  cd $cwd
  unset cwd
}
