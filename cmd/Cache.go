package cmd

import (
	"github.com/4color/SiSaiBlog/cmd/DbNav"
	"github.com/4color/SiSaiBlog/cmd/db"
)

//全局的参数
var GlobalBlogSeting *db.TBlogSetting

//导航数据
var GlobalNav *[]DbNav.TNav

func InitCache() {

	//获取表信息
	blog := db.TBlogSetting{}
	bloginfo, err := blog.GetInfo()
	if err != nil {
		GlobalBlogSeting = &blog
	} else {
		GlobalBlogSeting = &bloginfo
	}

	//获取导航
	modelNav := DbNav.TNav{}
	navs, err := modelNav.List("", 1, 100)
	if err != nil {
		datas := make([]DbNav.TNav, 0)
		GlobalNav = &datas
	} else
	{
		GlobalNav = &navs
	}
}
