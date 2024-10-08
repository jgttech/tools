# Development Scripts

A collection of software for CLI tools and utilities that I like.

## Prerequisites

- [Go](https://formulae.brew.sh/formula/go#default)
  - Tools are written in Go.
- [GVM](https://github.com/moovweb/gvm)
  - Helps with managing the Go version.
- [ZSH](https://github.com/ohmyzsh/ohmyzsh/wiki/Installing-ZSH)
  - The tools attempt to automatically register with your ZSH config (`${HOME}/.zshrc`).

## Setup

This will download the repo to a specific location on your system (default is `${HOME}/.tools`), attempt to compile the repo, and then run the freshly built binary install command. The install command will run the setup process and attempt to automate the installation the tools to be linked to your `${HOME}/.zshrc` file.

> **[WARNING]** <br /> This invokes a shell script for running the setup. You can review what the script does on GitHub. NEVER run a script blindly without making sure you trust it.

---

##### Variables

- `URL`
  - The URL to the shell file content. I used a short URL to keep things small.
- `DIR`
  - The name of the directory to install everything in, under the users `$HOME` path.

---

Pick one of these methods for installing the tools.

##### *Option 1 - wget*

```bash
URL="https://tinyurl.com/jgttech" DIR=".tools"; wget -qO- $URL | zsh -s -- $DIR;
```

##### *Option 2 - curl*

```bash
URL="https://tinyurl.com/jgttech" DIR=".tools"; curl -s $URL | zsh -s -- $DIR;
```
