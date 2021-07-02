package main

import (
	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	cmder "github.com/yaegashi/cobra-cmder"
)

type (
	App struct {
		Bool bool
	}

	AppSub1 struct {
		*App
		String string
	}

	AppSub1Sub2 struct {
		*AppSub1
		Int int
	}
)

func (app *App) Cmd() *cobra.Command {
	cmd := &cobra.Command{Use: "app"}
	cmd.PersistentFlags().BoolVarP(&app.Bool, "bool", "b", false, "input bool")
	return cmd
}

func (app *AppSub1) Cmd() *cobra.Command {
	cmd := &cobra.Command{Use: "sub1"}
	cmd.PersistentFlags().StringVarP(&app.String, "string", "s", "", "input string")
	return cmd
}

func (app *AppSub1Sub2) Cmd() *cobra.Command {
	cmd := &cobra.Command{Use: "sub2", Run: app.Run}
	cmd.Flags().IntVarP(&app.Int, "int", "i", 0, "input int")
	return cmd
}

func (app *AppSub1Sub2) Run(cmd *cobra.Command, args []string) {
	pp.Println(args)
	pp.Println(app.Int)
}

// 关联这步在做啥?

func (app *App) AppSub1xCommander() cmder.Cmder {
	return &AppSub1{App: app}
}

func (app *AppSub1) AppSub1Sub2Commander() cmder.Cmder {
	return &AppSub1Sub2{AppSub1: app}
}

func main() {
	app := &App{}
	cmd := cmder.Cmd(app)
	// cmd.SetArgs([]string{"sub1", "sub2", "-b", "-s", "abc", "-i", "123"})
	cmd.SetArgs([]string{"sub1", "sub2", "-i", "123"})
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
