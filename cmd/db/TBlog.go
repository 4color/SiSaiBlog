package db

import (
	"database/sql"
	"strconv"
	"strings"
)

type TBlog struct {
	Blogid      int64  `json:"blogid" form:"blogid"`
	Title       string `json:"title" form:"title"`
	Content     string `json:"content" form:"content"`
	Publishtime string `json:"publishtime" form:"publishtime"`
	Urltitle    string `json:"urltitle" form:"urltitle"`
	Author      string `json:"author" form:"author"`
	Source      string `json:"source" form:"source"`
	Tags        string `json:"tags" form:"tags"`
	Status      string `json:"status" form:"status"`
	Intro       string `json:"intro" form:"intro"`
	Views       int64  `json:"views" form:"views"`
	Comments    string `json:"comments" form:"comments"`
	Istop       int64  `json:"istop" form:"istop"`
	Smallpic    string `json:"smallpic" form:"smallpic"`
}

//保存
func (p *TBlog) Save(data TBlog) (set TBlog, err1 error) {

	set = data
	//如果包含了无行数据的消息，则不为错误
	if (data.Blogid == 0) {
		//新增
		count, err := SqlDB.Exec("insert into t_blogs (title,content,publishtime,urltitle,author,source,tags,views,comments,status,intro,istop,smallpic) "+
			"values (:1,:2,:3,:4,:5,:6,:7,:8,:9,:10,:11,:12,:13)",
			data.Title, data.Content, data.Publishtime, data.Urltitle, data.Author, data.Source, data.Tags, 0, 0, 1, data.Intro, data.Istop, data.Smallpic)
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
		_, err := SqlDB.Exec("update t_blogs set title =:1, content =:2, publishtime =:3,urltitle=:4,author=:5,source=:6,tags=:7,intro=:8,istop=:9,smallpic=:10 where blogid=:11",
			data.Title, data.Content, data.Publishtime, data.Urltitle, data.Author, data.Source, data.Tags, data.Intro, data.Istop, data.Smallpic, data.Blogid)
		if err != nil {
			err1 = err
			return
		}
	}

	return
}

//读取一条
func (p *TBlog) ReadOne(blogid string) (set TBlog, err error) {

	row := SqlDB.QueryRow("SELECT blogid,title,content,publishtime,urltitle,author,source,tags,intro,views,istop,smallpic FROM t_blogs where blogid=:1", blogid)
	err = row.Scan(&set.Blogid, &set.Title, &set.Content, &set.Publishtime, &set.Urltitle, &set.Author, &set.Source, &set.Tags, &set.Intro, &set.Views, &set.Istop, &set.Smallpic)

	//如果包含了无行数据的消息，则不为错误
	if (err != nil && strings.Contains(err.Error(), "no rows in result set")) {
		err = nil
	}
	return
}

//博客列表
func (p *TBlog) BlogList(keystring string, pageIndex int, pageSize int) (blogs []TBlog, err error) {

	blogs = make([]TBlog, 0)
	var sql = "SELECT blogid,title,publishtime,urltitle,author,source,tags,intro,comments,views,istop,smallpic " +
		" FROM t_blogs t WHERE title like  '%" + keystring + "%' and status = 1  order by istop desc, publishtime desc,blogid desc LIMIT {0},{1}"

	istart := (pageIndex - 1) * pageSize

	if (pageIndex > 1) {
		istart++
	}

	sql = strings.Replace(sql, "{0}", strconv.Itoa(istart), -1)
	sql = strings.Replace(sql, "{1}", strconv.Itoa(pageSize), -1)

	rows, err := SqlDB.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var set TBlog
		err = rows.Scan(&set.Blogid, &set.Title, &set.Publishtime, &set.Urltitle, &set.Author, &set.Source, &set.Tags, &set.Intro, &set.Comments, &set.Views, &set.Istop, &set.Smallpic)
		blogs = append(blogs, set)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

//博客列表总数
func (p *TBlog) BlogListCount(keystring string) (blogs int, err error) {
	blogs = 0
	row := SqlDB.QueryRow("SELECT COUNT(1) ZS FROM t_blogs t WHERE title like '%" + keystring + "%' and status = 1")
	row.Scan(&blogs)
	return
}

//博客列表总数
func (p *TBlog) BlogDelete(blogid string) (blogs int, err error) {

	_, err1 := SqlDB.Exec("update t_blogs set status=0 where blogid=:1", blogid)
	if err1 != nil {
		err = err1
		return
	}

	return
}

//更新查看数
func (p *TBlog) UpdateView(blogid string) (err1 error) {

	_, err := SqlDB.Exec("update t_blogs set views=views+1 where blogid=:1",
		blogid)
	if err != nil {
		err1 = err
	}
	return
}
