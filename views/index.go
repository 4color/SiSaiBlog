package views

import (
	"github.com/4color/SiSaiBlog/cmd/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Index(gc *gin.Context) {

	blog := db.TBlog{}
	pageindex := gc.DefaultQuery("p", "1")
	var pagesize = 10

	ipageindex, _ := strconv.Atoi(pageindex)
	keystring := gc.DefaultQuery("keystring", "")
	bloginfo, err := blog.BlogList(keystring, ipageindex, pagesize)
	if err != nil {
		bloginfo = make([]db.TBlog, 0)
	}

	var preindex = ipageindex - 1
	var nextindex = ipageindex + 1

	if (preindex == 0) {
		preindex = 1
	}
	gc.HTML(http.StatusOK, "index.html", gin.H{
		"blogs":     bloginfo,
		"preindex":  preindex,
		"nextindex": nextindex,
	})
}

//查看博客
func Viewblog(gc *gin.Context) {
	blogid := gc.Param("blogid")
	//获取表信息
	blog := db.TBlog{}

	bloginfo, err := blog.ReadOne(blogid)
	if err != nil {
		gc.HTML(http.StatusOK, "error.html", gin.H{
			"errormsg": err.Error(),
		})
	} else {
		gc.HTML(http.StatusOK, "blog.html", gin.H{
			"blog": bloginfo,
		})
	}
}
