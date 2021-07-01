package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tech-botao/go-implement/app/viper-cobra/cmd/sub"
)

var (
	Root *cobra.Command
)

func init() {
	Root = &cobra.Command{
		Use: "go-command",
		Short: "説明",
		RunE: func(c *cobra.Command, args []string) error {
			return nil
		},
	}

	Root.AddCommand(sub.CmdTestArgs, sub.CmdTestFlags, sub.CmdTestFile)
}
