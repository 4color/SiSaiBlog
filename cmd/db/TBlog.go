package db

import "database/sql"

type TBlog struct {
	Blogid      int64  `json:"blogid" form:"blogid"`
	Title       string `json:"title" form:"title"`
	Content     string `json:"content" form:"content"`
	Publishtime string `json:"publishtime" form:"publishtime"`
	Urltitle    string `json:"urltitle" form:"urltitle"`
	Author      string `json:"author" form:"author"`
	Source      string `json:"source" form:"source"`
	Tags        string `json:"tags" form:"tags"`
}

//保存
func (p *TBlog) Save(data TBlog) (set TBlog, err1 error) {

	set = data
	//如果包含了无行数据的消息，则不为错误
	if (data.Blogid == 0) {
		//新增
		count, err := SqlDB.Exec("insert into t_blogs (title,content,publishtime,urltitle,author,source,tags,views,comments) "+
			"values (:1,:2,:3,:4,:5,:6,:7,:8,:9)",
			data.Title, data.Content, data.Publishtime, data.Urltitle, data.Author, data.Source, data.Tags, 0, 0)
		if err != nil {

			err1 = err
			return
		} else {

			var result = count.(sql.Result)
			id, _ := result.LastInsertId()
			set.Blogid = id
		}

	} else
	{
		//update
		_, err := SqlDB.Exec("update t_blogs set title =:1, content =:2, publishtime =:3,urltitle=:4,author=:5,source=:6,tags=:7 where blogid=:8",
			data.Title, data.Content, data.Publishtime, data.Urltitle, data.Author, data.Source, data.Tags, data.Blogid)
		if err != nil {
			err1 = err
			return
		}
	}

	return
}
