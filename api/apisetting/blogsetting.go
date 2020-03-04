package apisetting

import (
	"github.com/4color/SiSaiBlog/api"
	"github.com/4color/SiSaiBlog/cmd/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSettingInfo(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	//获取表信息
	blog := db.TBlogSetting{}

	bloginfo, err := blog.GetInfo()
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	//返回值
	res.Data = bloginfo

	res.Status = http.StatusOK
	gc.JSON(200, res)
}

func SaveSettingInfo(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	var param db.TBlogSetting
	err := gc.BindJSON(&param)

	//获取表信息
	blog := db.TBlogSetting{}

	bloginfo, err := blog.Save(param)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}

	//返回值
	res.Data = bloginfo
	res.Message = "保存成功"

	res.Status = http.StatusOK
	gc.JSON(200, res)
}
