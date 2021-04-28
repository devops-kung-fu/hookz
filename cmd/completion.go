package cmd

import (
	"os"

	"github.com/devops-kung-fu/hookz/lib"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

	$ source <(yourprogram completion bash)

	# To load completions for each session, execute once:
	# Linux:
	$ yourprogram completion bash > /etc/bash_completion.d/yourprogram
	# macOS :
	$ yourprogram completion bash > /usr/local/etc/bash_completion.d/yourprogram

Zsh:

	# If shell completion is not already enabled in your environment,
	# you will need to enable it.  You can execute the following once:

	$ echo "autoload -U compinit; compinit" >> ~/.zshrc

	# To load completions for each session, execute once:
	$ yourprogram completion zsh > "${fpath[1]}/_yourprogram"

	# You will need to start a new shell for this setup to take effect.

fish:

	$ yourprogram completion fish | source

	# To load completions for each session, execute once:
	$ yourprogram completion fish > ~/.config/fish/completions/yourprogram.fish

PowerShell:

	PS> yourprogram completion powershell | Out-String | Invoke-Expression

	# To load completions for every new session, run:
	PS> yourprogram completion powershell > yourprogram.ps1
	# and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			lib.IfErrorLog(cmd.Root().GenBashCompletion(os.Stdout), "ERROR")
		case "zsh":
			lib.IfErrorLog(cmd.Root().GenZshCompletion(os.Stdout), "ERROR")
		case "fish":
			lib.IfErrorLog(cmd.Root().GenFishCompletion(os.Stdout, true), "ERROR")
		case "powershell":
			lib.IfErrorLog(cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout), "ERROR")
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
