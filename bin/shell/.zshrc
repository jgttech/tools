#!/usr/bin/env zsh
tools sync

# Tools
alias tools-dir="cd \${HOME}/$(tools env BASE_DIR)"

# Git
alias vim="nvim"
alias add="tools git add"
alias cm="tools git commit"
alias push="tools git push"
alias co="tools git co"
alias wip="tools git wip"

# Docker
alias ds="tools docker ls"
alias drc="tools docker remove containers"
alias dri="tools docker remove images"
alias drv="tools docker remove volumes"
alias drn="tools docker remove networks"
alias dn="drc; dri; drv; drn;"
