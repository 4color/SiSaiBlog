package adminview

import (
	"github.com/4color/SiSaiBlog/cmd"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func List(gc *gin.Context) {

	gc.HTML(http.StatusOK, "admin/nav.html", gin.H{
		"title":   "Posts",
		"blogset": cmd.GlobalBlogSeting,
	})
}

func Edit(gc *gin.Context) {

	navid := gc.DefaultQuery("navid", "0")
	inavid, _ := strconv.Atoi(navid)

	gc.HTML(http.StatusOK, "admin/navedit.html", gin.H{
		"navid":   inavid,
		"blogset": cmd.GlobalBlogSeting,
	})
}
