package cmd

import (
	"github.com/4color/SiSaiBlog/cmd/db"
)

//全局的参数
var GlobalBlogSeting *db.TBlogSetting

func InitCache() {

	//获取表信息
	blog := db.TBlogSetting{}
	bloginfo, err := blog.GetInfo()
	if err != nil {
		GlobalBlogSeting = &blog
	} else {
		GlobalBlogSeting = &bloginfo
	}
}
