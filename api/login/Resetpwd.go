package login

import (
	"github.com/4color/SiSaiBlog/api"
	"github.com/4color/SiSaiBlog/api/checkcode"
	"github.com/4color/SiSaiBlog/cmd/muser"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EditPwdString struct {
	Pwd         string `json:"pwd"`
	VerifyValue string `json:"VerifyValue"`
	CaptchaId   string `json:"captchaId"`
}

func EditPwd(gc *gin.Context) {
	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	var param EditPwdString
	err := gc.BindJSON(&param)

	//判断验证码
	var Codeparam checkcode.ConfigJsonBody
	Codeparam.VerifyValue = param.VerifyValue
	Codeparam.Id = param.CaptchaId
	codeResult := checkcode.CheckCodeVerify(Codeparam)

	if (codeResult["code"] == 0) {
		res.Message = codeResult["msg"].(string)
		gc.JSON(200, res)
		return
	}

	//获取表信息
	mUser := muser.Muser{}

	var username = api.GetSessionUserName(gc)

	err = mUser.EditPwd(username, param.Pwd)

	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)
		return
	}
	//返回值
	res.Message = "修改成功"

	res.Status = http.StatusOK
	gc.JSON(200, res)
}
