package sub

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
)

var CmdTestArgs *cobra.Command

func init() {
	CmdTestArgs = &cobra.Command{
		Use:                        "args",
		Short:                      "测试参数用法([]args)",
		Example:                    "Example: viper-cobra args 123 456 780",

		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 {
				return errors.New("no args")
			}
			fmt.Printf("args:%v", args)
			return nil
		},
	}
}
