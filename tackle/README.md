# Hookz Tacklebox

## Usage

Each ```yaml``` file contains an action that can be used anywhere in your ```.hookz.yaml``` file. Simply copy the non-commented lines and insert as an action under the hook type of your choice.

## Contributing

If you would like to contribute your own tackle, please follow the [contribution guidelines](./../CONTRIBUTING.md) and create a pull request. The pull request should have a yaml file with an action that can be copied and pasted into a .hookz.yaml file, and additions to the tackle section on this page to describe your action.

## Tackle

| Name                                             | Category | Action                                                       | Hook *           | Notes                                                        |
| ------------------------------------------------ | -------- | ------------------------------------------------------------ | ---------------- | ------------------------------------------------------------ |
| [errcheck.yaml](go/errcheck.yaml)                | Go       | Checks for errors that are not handled in go code            | ```pre-commit``` | Requires ```errcheck``` (https://github.com/kisielk/errcheck) |
| [errcheck-script.yaml](/go/errcheck-script.yaml) | Go       | Checks for errors that are not handled in go code. This is a script version that can be used to ignore errors | ```pre-commit``` | Requires ```errcheck``` (https://github.com/kisielk/errcheck) |
| [gocyclo.yaml](go/gocyclo.yaml)                  | Go       | Outputs cyclomatic complexity of all files.                  | ```pre-commit``` | Requires ```gocyclo```(https://github.com/fzipp/gocyclo)     |

\* The recommended hook type for this action