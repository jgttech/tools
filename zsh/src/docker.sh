##########################################
# Docker List (All)
function docker-list-all() {
  echo "\n[CONTAINERS]";
  docker ps -a;
  echo "\n";

  echo "[IMAGES]";
  docker images -a;
  echo "\n";

  echo "[VOLUMES]";
  docker volume ls;
  echo "\n";

  echo "[NETWORKS]";
  docker network ls;
  echo "\n";
}

##########################################
# Docker List Aliases
alias ds="docker-list-all";
alias dls="docker-list-all";

##########################################
# Docker Nuke Commands
function dk() {
  docker kill $(docker ps -q) &> /dev/null;
}

function drc() {
  docker rm -f $(docker ps -qa);
}

function dri() {
  docker rmi -f $(docker images -qq);
}

function drv() {
  docker volume rm -f $(docker volume ls -q);
}

function drn() {
  docker network rm -f $(docker network ls -q);
}

function dsp() {
  docker system prune -af
}

function dn() {
  dk;
  drc;
  dri;
  drn;
  drv;
  dsp;
}
