package apinav

import (
	"github.com/4color/SiSaiBlog/api"
	"github.com/4color/SiSaiBlog/cmd"
	"github.com/4color/SiSaiBlog/cmd/DbNav"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//保存
func Save(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	var param DbNav.TNav
	err := gc.BindJSON(&param)

	if (err != nil) {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	//获取表信息
	model := DbNav.TNav{}

	bloginfo, err := model.Save(param)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
	} else {
		//返回值
		res.Data = bloginfo
		res.Message = "保存成功"

		cmd.InitCache();
		res.Status = http.StatusOK
		gc.JSON(http.StatusOK, res)
	}
}

//编辑单篇博客
func Read(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	id := gc.Param("id")
	//获取表信息
	model := DbNav.TNav{}

	bloginfo, err := model.ReadOne(id)
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

//列表
func List(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	keystring := gc.Param("keystring")
	//获取表信息
	model := DbNav.TNav{}
	pageindex := gc.DefaultQuery("pageindex", "1")
	var pagesize = 10

	ipageindex, _ := strconv.Atoi(pageindex)
	bloginfo, err := model.List(keystring, ipageindex, pagesize)
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

func ListCount(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	keystring := gc.Param("keystring")
	//获取表信息
	model := DbNav.TNav{}

	bloginfo, err := model.ListCount(keystring)
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

func Delete(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	id := gc.DefaultQuery("navid", "0")
	//获取表信息
	model := DbNav.TNav{}

	bloginfo, err := model.Delete(id)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	//返回值
	res.Data = bloginfo
	res.Message = "删除成功"


	//更新缓存
	cmd.InitCache();

	res.Status = http.StatusOK
	gc.JSON(200, res)
}

