#!/usr/bin/env zsh
tools sync

# Tools
alias tools-dir="cd \${HOME}/$(tools env BASE_DIR)"

# Git
alias vim="nvim"
alias add="tools git add"
alias cm="tools git commit"
alias push="tools git push"
alias pull="tools git pull"
alias co="tools git co"
alias wip="tools git wip"
alias rebase="tools git rebase"

# Docker
alias ds="tools docker ls"
alias dn="tools docker nuke"
alias drc="tools docker remove containers"
alias dri="tools docker remove images"
alias drv="tools docker remove volumes"
alias drn="tools docker remove networks"
alias prune="tools docker prune"
