package blog

import (
	"github.com/4color/SiSaiBlog/api"
	"github.com/4color/SiSaiBlog/cmd/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//博客的编辑与保存
func SaveBlog(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	var param db.TBlog
	err := gc.BindJSON(&param)

	if (err != nil) {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	//获取表信息
	blog := db.TBlog{}

	bloginfo, err := blog.Save(param)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
	} else {
		//返回值
		res.Data = bloginfo
		res.Message = "保存成功"
		res.Status = http.StatusOK
		gc.JSON(http.StatusOK, res)
	}
}

//编辑单篇博客
func EditBlog(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	blogid := gc.Param("blogid")
	//获取表信息
	blog := db.TBlog{}

	bloginfo, err := blog.ReadOne(blogid)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	//返回值
	res.Data = bloginfo
	res.Message = "获取成功"

	res.Status = http.StatusOK
	gc.JSON(200, res)
}

//博客列表
func BlogList(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	keystring := gc.Param("keystring")
	//获取表信息
	blog := db.TBlog{}
	pageindex := gc.DefaultQuery("pageindex", "1")
	var pagesize = 10

	ipageindex, _ := strconv.Atoi(pageindex)
	bloginfo, err := blog.BlogList(keystring, ipageindex, pagesize)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}
	//返回值
	res.Data = bloginfo
	res.Message = "获取成功"

	res.Status = http.StatusOK
	gc.JSON(200, res)
}

func BlogListCount(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	keystring := gc.Param("keystring")
	//获取表信息
	blog := db.TBlog{}

	bloginfo, err := blog.BlogListCount(keystring)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}
	//返回值
	res.Data = bloginfo
	res.Message = "获取成功"

	res.Status = http.StatusOK
	gc.JSON(200, res)
}

func BlogDelete(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	blogid := gc.DefaultQuery("blogid", "0")
	//获取表信息
	blog := db.TBlog{}

	bloginfo, err := blog.BlogDelete(blogid)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}
	//返回值
	res.Data = bloginfo
	res.Message = "删除成功"

	res.Status = http.StatusOK
	gc.JSON(200, res)
}
