#!/usr/bin/env zsh
tools sync

alias tools-dir="cd \${HOME}/$(tools env BASE_DIR)"
alias vim="nvim"
alias add="tools git add"
alias cm="tools git commit"
alias push="tools git push"
alias co="tools git co"
alias wip="tools git wip"
alias ds="tools docker ls"
