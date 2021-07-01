package sub

import (
	"fmt"
	"github.com/spf13/cobra"
)

type (
	cmdTestFlagsArgs struct {
		Name string
		Age  int64
		Sex  bool
	}
)

var (
	CmdTestFlags      *cobra.Command
	_cmdTestFlagsArgs cmdTestFlagsArgs
)

func init() {
	CmdTestFlags = &cobra.Command{
		Use: "flags",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%v", _cmdTestFlagsArgs)
		},
	}
	_cmdTestFlagsArgs = cmdTestFlagsArgs{}
	CmdTestFlags.PersistentFlags().StringVar(&_cmdTestFlagsArgs.Name, "name", "who", "input you name")
	CmdTestFlags.PersistentFlags().Int64Var(&_cmdTestFlagsArgs.Age, "age", 99, "input you age")
	CmdTestFlags.PersistentFlags().BoolVar(&_cmdTestFlagsArgs.Sex, "sex", false, "input you sex")

}
