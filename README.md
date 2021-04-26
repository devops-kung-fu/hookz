

![](img/hookz-logo.png)

# hookz

[![Go Report Card](https://goreportcard.com/badge/github.com/devops-kung-fu/hookz)](https://goreportcard.com/report/github.com/devops-kung-fu/hookz) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/devops-kung-fu/hookz)

Manages git hooks inside a local git repository based on a configuration.

## Overview

Have you ever wanted to integrate custom functionality into your code commit and push workflows without using IDE plugins, hacks, or gnarly scripts? We have, and that's why we developed Hookz. Hookz allows us to do a ton of commit tasks _before_ our code even hits our branches. As you commit and push code to a git based source repository, Hookz will trigger scripts, run tasks, or do just about anything for you.

The best thing? We don't care what IDE you use to develop, what your back end source control system is, or what languages you program in. Operating System? We don't care. We wrote this in Go and have compiled it for every OS and architecture you'd ever want.

Here's what happens when we use Hookz on Hookz itself:

![](img/run-hookz.png)

## Installation

To install hookz,  [download the latest release](https://github.com/devops-kung-fu/hookz/releases) , make is executable, rename it to ```hookz``` and toss it in your ```/usr/local/bin``` directory for Linux, or on your path for other operating systems.

Linux Example:

```bash
sudo chmod +x hookz-2.1.1-linux-amd64
sudo mv hookz-2.1.1-linux-amd64 /usr/local/bin/hookz
```

## Configuration

Hookz uses a configuration file to generate hooks in your local git repository. This file needs to be in the root of your repository and must be named ```.hookz.yaml```

### Example Configuration

``` yaml
version: 2.1.1
hooks:
  - type: pre-commit
    actions:
      - name: "PlantUML Image Generator"
        url: https://github.com/jjimenez/pre-plantuml
        args: ["deflate"]
      - name: "Git Pull (Ensure there are no upstream changes)"
        exec: git
        args: ["pull"]
      - name: "Go Tidy"
        exec: go
        args: ["mod", "tidy"]
      - name: "Add all changed files during the pre-commit stage"
        exec: git
        args: ["add", "."]
  - type: post-commit
    actions:
    - name: "Post Echo"
      exec: echo
      args: ["-e", "Done!"]
```

Hooks will read this example configuration and create a ```pre-commit``` hook and a ```post-commit``` hook based on this yaml. 

An action with an ```URL``` will download the binary from the defined URL and configure the hook to execute the command with the defined arguments before a commit happens.

The post-commit in this configuration will execute a command named "dude" with the arguments "Hello World" after a commit has occurred. Note that the _dude_ command must be on your path. If it isn't this post-commit will fail because the command isn't found.

### Optional elements

The following notes apply to the elements in the YAML:

|Attribute|Notes|
|---|---|
|```URL```|If this exists, then exec and script are ignored. The URL must be a link to an executable binary|
|```exec```|If this exists then URL and script are ignored|
|```script```|If this exists then URL, exec, and args are ignored|
|```args```|Optional in all cases|

### Inline scripting

Scripts can be embedded into the ```.hookz.yaml``` in multiline format such as follows:

__NOTE:__ There needs to be a \n at the end of a line if a multi-line statement exists in the ```script``` node, and special characters need to be escaped properly. 

``` yaml 
- type: pre-commit
    actions:
      - name: "Go Tidy (Recursive)"
        script: "
          #!/bin/bash \n
          echo -e Tidying all found go.mod occurrences \n
          find . -name go.mod -print0 | xargs -0 -n1 dirname |  xargs -L 1 bash -c 'cd \"$0\" && pwd && go mod tidy' \n
          "
```
If you have args flags set, they can be referenced as $1, $2, etc. in your script in a similar manner as passing parameters in. Any scripting language is supported.

### Support for multiple commands in a hook

If multiple hooks are defined in the configuration with the same type (ie: ```pre-commit```) they will be configured to run in the order they appear in the file. There is no need to group types together, they will be written to the appropriate hooks.

### Hook types

Hook types that will execute are the same as supported by _git_. Examples are as follows:

* ```applypatch-msg```
* ```commit-msg```
* ```fsmonitor-watchman```
* ```post-commit```
* ```post-update```
* ```pre-applypatch```
* ```pre-commit```
* ```pre-update```
* ```prepare-commit-msg```
* ```pre-push```
* ```pre-rebase```
* ```pre-receive```
* ```update```

### Return Codes

Any non-zero return code from a command executed in a hook will return a FAIL.

### Security

**IMPORTANT:** the ```URL``` element will download a binary, store it in your .git/hookz folder and mark it executable. It is important to ensure that **YOU KNOW THE SAFETY** of the executable that you are downloading. A more secure way is to use the ```exec``` attribute to run an application which is already locally installed, or embed a script in the ```script``` element to download and check the ```shasum``` of the file.

For additional security notes, view the [SECURITY.md](SECURITY.md)

## Running Hookz

![](img/hookz.png)

To generate the hooks as defined in your configuration simply execute the following command in the _root of your local repository_ where the ```.hookz.yaml``` file resides:

``` bash
hookz initialize # you can also use the init alias
```

Removing hooks can be done by executing the following command:

``` bash
hookz remove
```

To re-download any file defined in an URL key:

``` bash
hookz update
```

### Applying changes to the .hookz.yaml
If there is a modification to the ```.hookz.yaml``` file in your application, you'll need to apply the changes using the following:

``` bash
hookz reset
```

### Verbose option

The ```initialize``` (init) and ```reset``` command optionally take a verbosity flag to indicate extended output should be displayed when a hook executes. This is handy for debugging or seeing errors that may be suppressed by hookz.

```
hookz init --verbose
hookz reset --verbose
```


## Example Hooks

### Recursively tidy all go.mod files in subdirectories

```yaml
version: 2.1.1
hooks:
  - type: pre-commit
    actions:
      - name: "Go Tidy (Recursive)"
        script: "
          #!/bin/bash \n
          echo -e Tidying all found go.mod occurrences\n
          find . -name go.mod -print0 | xargs -0 -n1 dirname |  xargs -L 1 bash -c 'cd \"$0\" && pwd && go mod tidy' \n
          "
```
### Update all go modules to the latest version before committing

```yaml
version: 2.1.1
hooks:
  - type: pre-commit
    actions:
      - name: "Update all go dependencies to latest"
        exec: go
        args: ["get", "-u", "./..."]
```

### Pull from your remote branch before committing

``` yaml
version: 2.1.1
hooks:
  - type: pre-commit
    actions:
      - name: "Git Pre-Commit Pull"
        exec: git
        args: ["pull"]
```
### Check for open source component vulnerabilities
Requires [Sonatype Nancy](https://ossindex.sonatype.org/integration/nancy)

__NOTE:__ The ```|| true``` on the end of the command makes the hook always return a success return code. If you want to break the commit on a failure, then remove this directive.
 
```yaml
version: 2.1.1
hooks:
  - type: pre-commit
    actions:
      - name: Sonatype Nancy
        script: "
          #!/bin/bash \n
          go list -json -m all | nancy sleuth || true \n
        "
```

### Terraform Format and Terraform Docs

Assumes `terraform` is in your `PATH` for `fmt`. 

```yaml
version: 2.1
hooks:
- type: pre-commit
  actions:
    - name: Terraform Format
      exec: terraform
      args: ["fmt"]
    - name: Terraform Docs
      url: https://github.com/terraform-docs/terraform-docs/releases/download/v0.12.1/terraform-docs-v0.12.1-linux-amd64
      args: ["markdown", "table", "--output-file", "README.md", "."]
    - name: "Add all changed files during the pre-commit stage"
      exec: git
      args: ["add", "."]
```

`README.md` must contain the following tags where the documentation will be injected.

```html
<!-- BEGIN_TF_DOCS -->

<!-- END_TF_DOCS -->
```

### NPM

```yaml
version: 2.1
hooks:
- type: pre-commit
  actions:
    - name: NPM CI
      exec: npm
      args: ["ci"]
    - name: NPM Test
      exec: npm
      args: ["test"]
    - name: "Add all changed files during the pre-commit stage"
      exec: git
      args: ["add", "."]
```

## Credits

A big thank-you to our friends at [Freepik](https://www.freepik.com) for the Hookz logo.