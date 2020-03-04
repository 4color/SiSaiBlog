package blog

import (
	"github.com/4color/SiSaiBlog/api"
	"github.com/4color/SiSaiBlog/cmd/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

//博客的编辑与保存
func EditBlog(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	var param db.TBlog
	err := gc.BindJSON(&param)

	//获取表信息
	blog := db.TBlog{}

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
