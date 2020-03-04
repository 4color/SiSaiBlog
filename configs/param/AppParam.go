package param

import (
	"fmt"
	"github.com/spf13/viper"
)

type Server struct {
	Port string
	Name string
}

type Db struct {
	ConnectionString string
}

type auth struct {
	Enable bool
	Authcode string
}

type AppParam struct {
	Mode string
	Server
	Db

}

func (p *AppParam)ReadConfig() {

	//default config
	viper.SetDefault("application.mode", "debug")
	viper.SetDefault("application.server.port", "8080")
	viper.SetDefault("application.server.name", "skdp-admin-service")

	//read properties
	p.Mode = viper.GetString("application.mode")
	//注意server的port数字前面有个冒号
	p.Server.Port = ":" + viper.GetString("application.server.port")
	p.Server.Name = viper.GetString("application.server.name")

	//数据库连接
	p.Db.ConnectionString=  viper.GetString("application.db.ConnectionString")


	fmt.Println("[Init AppParam Over]...", p)

}