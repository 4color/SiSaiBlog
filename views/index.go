package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(gc *gin.Context) {
	gc.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Posts",
	})
}

