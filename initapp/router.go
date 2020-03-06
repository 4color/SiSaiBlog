package initapp

import (
	"github.com/4color/SiSaiBlog/api"
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

	//自定义分割符
	R.Delims("{[{", "}]}")
	//R.SetFuncMap(template.FuncMap{
	//	"ToHtml1": ToHtml1,
	//})
	//R.SetFuncMap(template.FuncMap{
	//	"ToHtml2": ToHtml2,
	//})

	R.Static("/static", "web/static")
	R.LoadHTMLGlob("web/template/**")

	R.GET("/", views.Index)
	R.GET("/view/*blogurl", views.Viewblog)


	store := cookie.NewStore([]byte("secret"))
	R.Use(sessions.Sessions("mysession", store))

	authorized := R.Group("/", api.AuthSessionMiddle())
	{

		authorized.GET("/admin", views.Admin)
		authorized.GET("/adminblog", views.AdminBlog)
		authorized.GET("/admin/blog", views.AdminNew)
		authorized.GET("/admin/reset", views.AdminReset)
	}

	R.GET("/login", views.Login)

	R.GET("/admin/logout", views.AdminLogout)

	v1 := R.Group("/api/v1", )
	{
		//验证码
		v1.GET("/checkcode", checkcode.Checkcode)
		v1.POST("/login", login.Login)

		//基本信息
		v1.GET("/blogsetting", apisetting.GetSettingInfo)
		//博客
		v1.GET("/blog/:blogid", blog.EditBlog)
		v1.GET("/bloglist", blog.BlogList)
		v1.GET("/blogs/count", blog.BlogListCount)
		v1.POST("/blogs/updateview", blog.BlogUpdateViews)

		authorizedapi := v1.Group("/", api.AuthSessionMiddle())
		{
			authorizedapi.POST("/blogs/delete", blog.BlogDelete)
			authorizedapi.POST("/blog", blog.SaveBlog)
			authorizedapi.POST("/blogsetting", apisetting.SaveSettingInfo)
			authorizedapi.POST("/repwd", login.EditPwd)
			authorizedapi.POST("/logout", login.Logout)
		}

	}
}
