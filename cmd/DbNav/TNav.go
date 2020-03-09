package DbNav

import (
	"database/sql"
	"github.com/4color/SiSaiBlog/cmd/db"
	"strconv"
	"strings"
)

type TNav struct {
	Navid    int64  `json:"navid" form:"navid"`
	Navcode  string `json:"navcode" form:"navcode"`
	Navname  string `json:"navname" form:"navname"`
	Link     string `json:"link" form:"link"`
	Target   string `json:"target" form:"target"`
	Navorder int64  `json:"navorder" form:"navorder"`
}

//保存
func (p *TNav) Save(data TNav) (set TNav, err1 error) {

	set = data
	//如果包含了无行数据的消息，则不为错误
	if (data.Navid == 0) {
		//新增
		count, err := db.SqlDB.Exec("insert into t_nav (navcode,navname,link,target,navorder) "+
			"values (:1,:2,:3,:4,:5)",
			data.Navcode, data.Navname, data.Link, data.Target, data.Navorder)
		if err != nil {
			err1 = err
			return
		} else {
			var result = count.(sql.Result)
			id, _ := result.LastInsertId()
			set.Navid = id
		}

	} else
	{
		//update
		_, err := db.SqlDB.Exec("update t_nav set navcode =:1, navname =:2, link =:3,target=:4,navorder=:5 where navid=:6",
			data.Navcode, data.Navname, data.Link, data.Target, data.Navorder, data.Navid)
		if err != nil {
			err1 = err
			return
		}
	}

	return
}

//读取一条
func (p *TNav) ReadOne(navid string) (data TNav, err error) {

	row := db.SqlDB.QueryRow("SELECT navid,navcode,navname,link,target,navorder FROM t_nav where navid=:1", navid)
	err = row.Scan(&data.Navid, &data.Navcode, &data.Navname, &data.Link, &data.Target, &data.Navorder)

	//如果包含了无行数据的消息，则不为错误
	if (err != nil && strings.Contains(err.Error(), "no rows in result set")) {
		err = nil
	}
	return
}

//列表
func (p *TNav) List(keystring string, pageIndex int, pageSize int) (datas []TNav, err error) {

	datas = make([]TNav, 0)
	var sql = "SELECT navid,navcode,navname,link,target,navorder " +
		" FROM t_nav  order by navorder asc"

	istart := (pageIndex - 1) * pageSize

	if (pageIndex > 1) {
		istart++
	}

	sql = strings.Replace(sql, "{0}", strconv.Itoa(istart), -1)
	sql = strings.Replace(sql, "{1}", strconv.Itoa(pageSize), -1)

	rows, err := db.SqlDB.Query(sql)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var data TNav
		err = rows.Scan(&data.Navid, &data.Navcode, &data.Navname, &data.Link, &data.Target, &data.Navorder)
		datas = append(datas, data)
	}
	if err = rows.Err(); err != nil {
		return
	}
	return
}

//列表总数
func (p *TNav) ListCount(keystring string) (count int, err error) {
	count = 0
	row := db.SqlDB.QueryRow("SELECT COUNT(1) ZS FROM t_nav")
	row.Scan(&count)
	return
}

func (p *TNav) Delete(navid string) (blogs int, err error) {

	_, err1 := db.SqlDB.Exec("delete from t_nav where navid=:1", navid)
	if err1 != nil {
		err = err1
		return
	}

	return
}
