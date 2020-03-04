package main

import (
	"github.com/4color/SiSaiBlog/configs"
	"github.com/4color/SiSaiBlog/initapp"
)

func main() {
	initapp.InitRouter()
	initapp.R.Run(configs.Config.App.Server.Port)
}