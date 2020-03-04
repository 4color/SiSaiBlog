package db

import (
	"strings"
)

type TBlogSetting struct {
	Systemid int64  `json:"systemid" form:"systemid"`
	Blogname string `json:"blogname" form:"blogname"`
	Github   string `json:"github" form:"github"`
	Counter  string `json:"counter" form:"counter"`
}

//获取信息
func (p *TBlogSetting) GetInfo() (set TBlogSetting, err error) {

	row := SqlDB.QueryRow("SELECT systemid,blogname,github,counter FROM t_blogsetting LIMIT 1")
	err = row.Scan(&set.Systemid, &set.Blogname, &set.Github, &set.Counter)

	//如果包含了无行数据的消息，则不为错误
	if (err != nil && strings.Contains(err.Error(), "no rows in result set")) {
		err = nil
	}

	return
}

//保存
func (p *TBlogSetting) Save(settingdata TBlogSetting) (set TBlogSetting, err error) {

	row := SqlDB.QueryRow("SELECT systemid,blogname,github,counter FROM t_blogsetting LIMIT 1")
	err = row.Scan(&set.Systemid, &set.Blogname, &set.Github, &set.Counter)

	//如果包含了无行数据的消息，则不为错误
	if (err != nil && strings.Contains(err.Error(), "no rows in result set")) {
		//新增
		_, err1 := SqlDB.Exec("insert into t_blogsetting (blogname,github,counter) "+
			"values (:1,:2,:3)",
			settingdata.Blogname, settingdata.Github, settingdata.Counter)
		if err1 != nil {
			return
		}

	} else
	{
		//update
		_, err1 := SqlDB.Exec("update t_blogsetting set blogname =:1, github =:2,counter=:3",
			settingdata.Blogname, settingdata.Github, settingdata.Counter)
		if err1 != nil {
			return
		}
	}

	return
}
