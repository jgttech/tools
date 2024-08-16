# Development Scripts

A collection of software for CLI tools and utilities that I like.

## Setup

> **[WARNING]** <br /> This invokes a shell script for running the setup. You can review what the script does on GitHub. NEVER run a script blindly without making sure you trust it.

Pick one of these methods for installing the tools.

##### *Option 1 - wget*

```bash
URL="https://shorturl.at/VI7oq"; wget -qO- "${URL}?$(date +%s)" | sh; unset URL;
```

##### *Option 2 - curl*

```bash
URL="https://shorturl.at/VI7oq"; curl -s "${URL}?$(date +%s)" | sh; unset URL;
```
