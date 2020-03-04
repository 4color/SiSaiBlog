package initapp

import (
	"github.com/4color/SiSaiBlog/api/apisetting"
	"github.com/4color/SiSaiBlog/api/blog"
	"github.com/4color/SiSaiBlog/api/checkcode"
	"github.com/4color/SiSaiBlog/api/login"
	"github.com/4color/SiSaiBlog/views"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var R *gin.Engine

func InitRouter() {
	R = gin.Default()

	store := cookie.NewStore([]byte("secret"))
	R.Use(sessions.Sessions("mysession", store))

	//自定义分割符
	R.Delims("{[{", "}]}")

	R.Static("/static", "web/static")
	R.LoadHTMLGlob("web/template/**")

	R.GET("/", views.Index)
	R.GET("/admin", views.Admin)

	R.GET("/login", views.Login)

	R.GET("/adminblog", views.AdminBlog)

	R.GET("/admin/blog", views.AdminNew)

	R.GET("/admin/reset", views.AdminReset)

	R.GET("/admin/logout", views.AdminLogout)

	v1 := R.Group("/api/v1")
	{
		//验证码
		v1.GET("/checkcode", checkcode.Checkcode)
		v1.POST("/login", login.Login)
		v1.POST("/repwd", login.EditPwd)
		v1.POST("/logout", login.Logout)

		//基本信息
		v1.GET("/blogsetting", apisetting.GetSettingInfo)
		v1.POST("/blogsetting", apisetting.SaveSettingInfo)

		//博客
		v1.POST("/blog", blog.EditBlog)

	}
}
