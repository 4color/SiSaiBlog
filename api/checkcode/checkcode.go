package checkcode

import (
	"github.com/4color/SiSaiBlog/api"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"net/http"
)

//获取验证码图片的验证base64码
//利用验证码id对比是不是某个客户端的验证码
//后期的验证要存在Redis,才可以实现扩展。
//参考 https://www.jianshu.com/p/08319bdd0637
//技术: github.com/mojocn/base64Captcha

type ConfigJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

//获取验证码
func Checkcode(gc *gin.Context) {

	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}

	param := getCodeString()
	var driver base64Captcha.Driver

	//create base64 encoding captcha
	switch param.CaptchaType {
	case "audio":
		driver = param.DriverAudio
	case "string":
		driver = param.DriverString.ConvertFonts()
	case "math":
		driver = param.DriverMath.ConvertFonts()
	case "chinese":
		driver = param.DriverChinese.ConvertFonts()
	default:
		driver = param.DriverDigit
	}
	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := c.Generate()
	body := map[string]interface{}{"code": 1, "data": b64s, "captchaId": id, "msg": "success"}
	if err != nil {
		body = map[string]interface{}{"code": 0, "msg": err.Error()}
	}

	res.Status = http.StatusOK
	res.Data = body

	gc.JSON(200, res)
}

//验证证码是否正确
func CheckCodeVerifyHandle(gc *gin.Context) {

	res := api.ResponseBodyModel{http.StatusInternalServerError, "", ""}
	var param ConfigJsonBody
	err := gc.BindJSON(&param)
	if err != nil {
		res.Message = err.Error()
		gc.JSON(http.StatusOK, res)

	}
	//verify the captcha
	body := CheckCodeVerify(param)

	res.Data = body;
	gc.JSON(http.StatusOK, res)
}

func CheckCodeVerify(param ConfigJsonBody) (result map[string]interface{}) {

	//verify the captcha
	result = map[string]interface{}{"code": 0, "msg": "验证码不正确"}
	if store.Verify(param.Id, param.VerifyValue, true) {
		result = map[string]interface{}{"code": 1, "msg": "ok"}
	}
	return
}

func getCodeString() (confg ConfigJsonBody) {

	confg.CaptchaType = "string"
	var da = base64Captcha.DriverString{
		Height:          60,
		Width:           240,
		ShowLineOptions: 0,
		NoiseCount:      0,
		Source:          "1234567890qwertyuioplkjhgfdsazxcvbnm",
		Length:          4,
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	confg.DriverString = &da

	return

}
