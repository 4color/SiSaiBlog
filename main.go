package main

import (
	"github.com/4color/SiSaiBlog/cmd"
	"github.com/4color/SiSaiBlog/configs"
	"github.com/4color/SiSaiBlog/initapp"
)

func main() {
	initapp.InitRouter()
	//实始化缓存
	cmd.InitCache()

	initapp.R.Run(configs.Config.App.Server.Port)
}
