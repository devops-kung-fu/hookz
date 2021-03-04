# hookz

Manages commit hooks inside a local git repository based on a configuration.


## Configuration

Hookz uses a configuration file to generate hooks in your local git repository. This file needs to be in the root of your repository and must be named *.hooks.yaml*

Take for example the following configuration:

``` yaml
  hooks:
  - name: "PlantUML Image Generator"
    type: pre-commit
    url: https://github.com/jjimenez/pre-plantuml
    args: [deflate]
  - name: "Post-Commit Echo"
    type: post-commit
    exec: derp
    args: ["Hello World"]

```

Hooks will read this configuration and create a pre-commit hook and a post-commit hook based on this yaml. 

The pre-commit will download the binary from the defined URL and configure the pre-commit to execute the command with the defined arguments before a commit happens.

The post-commit in this configuration will execute a command named "derp" with the arguments "Hello World" after a commit has occurred. Note that the _derp_ command must be on your path. If it isn't this post-commit will fail because the command isn't found.

## Running Hookz

To generate the hooks as defined in your configuration simply execute the following:

``` bash
hookz init
```

Removing hooks can be done by executing the following command:

``` bash
hookz remove
```
## Example Hooks

### Pull from your remote branch before committing

``` yaml
  - name: "Git Pre-Commit Pull"
    type: pre-commit
    exec: git
    args: [pull]
```



