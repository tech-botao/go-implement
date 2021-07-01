package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/k0kubun/pp"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

type (
	Config struct {
		Bar struct {
			A int    `json:"a"`
			B int    `json:"b"`
			C string `json:"c"`
			D int    `json:"d"`
		} `json:"bar"`
		Baz struct {
			Awesome string `json:"awesome"`
			Bad     string `json:"bad"`
			Great   string `json:"great"`
		} `json:"baz"`
		Foo struct {
			A int `json:"a"`
			B int `json:"b"`
			C int `json:"c"`
		} `json:"foo"`
		ID     string `json:"id"`
		Mid    int    `json:"mid"`
		Status struct {
			Awesome string `json:"awesome"`
			Work    bool   `json:"work"`
		} `json:"status"`
	}
)

func init() {
	// 初始化的时候, 可以设置各种信息
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// 多个文件目录进行检索
	viper.AddConfigPath("/etc/go/")
	viper.AddConfigPath("$HOME/.appname")
	viper.AddConfigPath("./app/viper-config")
	viper.AddConfigPath(".")

	// 读取文件
	err := viper.ReadInConfig()               // Find and read the config file
	if err != nil {                           // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	// 变更监控
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		pp.Println("Config file changed:", e.Name)
		pp.Println("config.foo", viper.AllSettings())
	})
}

func main() {

	// 读取相关信息
	pp.Println("config", viper.AllKeys())
	pp.Println("config.foo", viper.AllSettings())

	// env
	viper.SetEnvPrefix("go")
	viper.BindEnv("id")
	os.Setenv("GO_ID", "12345")
	pp.Println(viper.Get("id"))

	// flag
	pflag.Int("mid", 1234, "message id")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	pp.Println(viper.GetInt("mid"))

	// Env优先, 如果有同名变量, 直接更新设定值
	viper.AutomaticEnv()

	// 输出另一个文件
	err := viper.WriteConfigAs("./app/viper-config/out.json")
	if err != nil {
		panic(err)
	}

	// 多个文件读取
	x := viper.NewWithOptions()
	x.SetConfigName("out")
	x.SetConfigType("json")
	x.AddConfigPath("./app/viper-config")
	x.ReadInConfig()
	var c Config
	x.Unmarshal(&c)
	pp.Println(c)

	// etcd
	//viper.AddRemoteProvider("etcd", "http://127.0.0.1:4001","/config/hugo.json")
	//viper.SetConfigType("json") // because there is no file extension in a stream of bytes, supported extensions are "json", "toml", "yaml", "yml", "properties", "props", "prop", "env", "dotenv"
	//err := viper.ReadRemoteConfig()

	// consul
	//viper.AddRemoteProvider("consul", "localhost:8500", "MY_CONSUL_KEY")
	//viper.SetConfigType("json") // Need to explicitly set this to json
	//err := viper.ReadRemoteConfig()
	//fmt.Println(viper.Get("port")) // 8080
	//fmt.Println(viper.Get("hostname")) // myhostname.com

	// 多个viper
	//x := viper.New()
	//y := viper.New()
	//x.SetDefault("ContentDir", "content")
	//y.SetDefault("ContentDir", "foobar")

	// 待机
	forever := make(chan bool)
	<-forever
}
