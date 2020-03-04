package configs

import (
	"fmt"
	p "github.com/4color/SiSaiBlog/configs/param"
	"github.com/spf13/viper"
)

type ConfigParam struct {
	App p.AppParam    //应用参数
}
var Config ConfigParam

func init() {


	fmt.Println("[初始化参数]......")

	viper.SetConfigType("yaml")
	viper.SetConfigName("application")   // name of config file (without extension)
	viper.AddConfigPath("./configs/yaml") // 配置文件路径，多次使用可以查找多个目录
	err := viper.ReadInConfig()          // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}

	//加载配置
	Config.App.ReadConfig()

	fmt.Println("[结束初始化参数]......")
}


