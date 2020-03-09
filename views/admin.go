package views

import (
	"github.com/4color/SiSaiBlog/api"
	"github.com/4color/SiSaiBlog/cmd"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Admin(gc *gin.Context) {

	gc.HTML(http.StatusOK, "admin/admin.html", gin.H{
		"title": "Posts",
		"blogset":   cmd.GlobalBlogSeting,
	})
}

func AdminListBlog(gc *gin.Context) {
	gc.HTML(http.StatusOK, "admin/list.html", gin.H{
		"title": "Posts",
		"blogset":   cmd.GlobalBlogSeting,
	})
}

func AdminNew(gc *gin.Context) {

	blogid := gc.DefaultQuery("blogid","0")
	iblogid, _ := strconv.Atoi(blogid)
	gc.HTML(http.StatusOK, "admin/editblog.html", gin.H{
		"blogid": iblogid,
		"blogset":   cmd.GlobalBlogSeting,
	})
}

func AdminReset(gc *gin.Context) {
	gc.HTML(http.StatusOK, "admin/reset.html", gin.H{
		"title": "Posts",
		"blogset":   cmd.GlobalBlogSeting,
	})
}

func AdminLogout(gc *gin.Context) {

	api.ClearSession(gc)
	gc.HTML(http.StatusOK, "admin/logout.html", gin.H{
		"title": "Posts",
		"blogset":   cmd.GlobalBlogSeting,
	})
}

