package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveSession(gc *gin.Context, tokenString string, username string) {

	session := sessions.Default(gc)
	session.Set("sisaiblog", tokenString)
	session.Set("sisaiblog_username", username)
	session.Save()
}

func GetSessionUserName(gc *gin.Context) (username string) {

	session := sessions.Default(gc)
	value := session.Get("sisaiblog_username")
	if (value != nil) {
		username = value.(string)
	} else {
		username = ""
	}
	return
}

func ClearSession(gc *gin.Context) {

	session := sessions.Default(gc)
	session.Clear()
	session.Save()
}


//认证
func AuthSessionMiddle() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionValue := session.Get("sisaiblog_username")
		if sessionValue == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorized",
			})
			c.Abort()
			return
		}
		// 设置简单的变量
		c.Set("sisaiblog_username", sessionValue.(string))

		c.Next()
		return
	}
}
