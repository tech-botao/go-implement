package sub

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Name string
		Age int64
	}

)

var (
	fileName string
	CmdTestFile *cobra.Command

	config Config
)

func init() {
	CmdTestFile = &cobra.Command{
		Use: "file",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%v", config)
		},
	}
	CmdTestFile.PersistentFlags().StringVarP(&fileName, "config", "c", "config.yaml", "config file name")

	cobra.OnInitialize(func() {
		v := viper.New()
		v.SetConfigType("yaml")
		v.SetConfigName(fileName)
		v.AddConfigPath(".")

		err := v.ReadInConfig()
		if err != nil {
			panic(err)
		}

		err = v.Unmarshal(&config)
		if err != nil {
			panic(err)
		}
	})
}


