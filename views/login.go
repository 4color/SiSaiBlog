package views

import (
	"github.com/4color/SiSaiBlog/cmd"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(gc *gin.Context) {
	gc.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Posts",
		"blogset":   cmd.GlobalBlogSeting,
	})
}

