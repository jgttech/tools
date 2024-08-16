# alias add="git add";
# alias cm="git commit -m";
# alias pull="git pull";
# alias push="git push";
# alias clone="gh repo clone";
# alias status="git status";
# alias co="git fetch; git checkout";
# alias fetch="git fetch";

function signer {
  export GPG_TTY=$(tty)
}

function add {
  signer;
  git add $@;
}

function cm {
  signer;
  git commit -m $@;
}

function cs {
  signer;
  git commit -S -m $@;
}

function pull {
  signer;
  git pull $@;
}

function push {
  signer;
  git push $@;
}

function clone {
  signer;
  git clone $@;
}

function status {
  signer;
  git status $@;
}

function co {
  signer;
  git fetch;
  git checkout $@;
}

function fetch {
  signer;
  git fetch $@;
}

function branch {
  git checkout -b $1;
  git push --set-upstream origin $1;
}

function rebase {
  branch=${1:-"main"}
  git rebase $branch
}

function merge {
  CURRENT_GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD);

  echo "Merging from '$1' to '${CURRENT_GIT_BRANCH}'";

  git checkout $1;
  git pull;
  git checkout ${CURRENT_GIT_BRANCH};

  unset CURRENT_GIT_BRANCH;

  git merge $1;
}

function wip {
  timestamp=$(date +%F\ %T)

  signer;
  git fetch;
  git add .;
  git commit -m "WIP ${timestamp}";
  git push;
}

function swip {
  timestamp=$(date +%F\ %T)

  signer;
  git fetch;
  git add .;
  git commit -S -m "WIP ${timestamp}";
  git push;
}
