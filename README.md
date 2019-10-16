## Description

A lightweight CLI command written in GoLang to interact with most of the useful Azure Devops resources, such as builds, deployments and work items.

## How to

After running `go build`, executable file `devops` is made

1. Copy the executable to your PATH folder, for example: `cp devops /usr/local/bin`
2. Move to your DevOps project folder
3. Start with `devops setup` to create a configuration file. Then `devops builds` or `devops deployments` commands will be available