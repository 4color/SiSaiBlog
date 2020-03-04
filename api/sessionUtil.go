package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
	username = value.(string)
	return
}

func ClearSession(gc *gin.Context) {

	session := sessions.Default(gc)
	session.Clear()
}