package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(gc *gin.Context) {
	gc.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Posts",
	})
}

