package muser

import (
	"database/sql"
	"github.com/4color/SiSaiBlog/cmd/db"
	"strings"
)

type Muser struct {
	Userid        string         `json:"userid" form:"userid"`
	Username      string         `json:"username" form:"username"`
	Loginip       sql.NullString `json:"loginip" form:"loginip"`
	Logintime     sql.NullString `json:"logintime" form:"logintime"`
	Lastloginip   sql.NullString `json:"lastloginip" form:"lastloginip"`
	Lastlogintime sql.NullString `json:"lastlogintime" form:"lastlogintime"`
}

//用户登录
func (p *Muser) GetUserLogin(username string, pwd string) (user Muser, err error) {

	row := db.SqlDB.QueryRow("SELECT username,userid,logintime,loginip,lastlogintime,lastloginip FROM t_user where username=? and password=? ", username, pwd)
	err = row.Scan(&user.Username, &user.Userid, &user.Logintime, &user.Loginip, &user.Lastlogintime, &user.Lastloginip)

	//如果包含了无行数据的消息，则不为错误
	if (err != nil && strings.Contains(err.Error(), "no rows in result set")) {
		err = nil
	}

	return
}

//修改密码
func (p *Muser) EditPwd(username string, pwd string) (err1 error) {

	_, err := db.SqlDB.Exec("update t_user set password=:1 where username=:2",
		pwd, username)
	//如果包含了无行数据的消息，则不为错误
	if (err != nil && strings.Contains(err.Error(), "no rows in result set")) {
		err1 = nil
	}

	return
}
