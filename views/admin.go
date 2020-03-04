package views

import (
	"github.com/4color/SiSaiBlog/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Admin(gc *gin.Context) {
	gc.HTML(http.StatusOK, "admin.html", gin.H{
		"title": "Posts",
	})
}

func AdminBlog(gc *gin.Context) {
	gc.HTML(http.StatusOK, "adminlist.html", gin.H{
		"title": "Posts",
	})
}

func AdminNew(gc *gin.Context) {
	gc.HTML(http.StatusOK, "editblog.html", gin.H{
		"title": "Posts",
	})
}

func AdminReset(gc *gin.Context) {
	gc.HTML(http.StatusOK, "adminreset.html", gin.H{
		"title": "Posts",
	})
}

func AdminLogout(gc *gin.Context) {

	api.ClearSession(gc)
	gc.HTML(http.StatusOK, "adminlogout.html", gin.H{
		"title": "Posts",
	})
}
