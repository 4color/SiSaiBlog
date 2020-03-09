package initapp

import (
	"github.com/4color/SiSaiBlog/api"
	"github.com/4color/SiSaiBlog/api/apinav"
	"github.com/4color/SiSaiBlog/api/apisetting"
	"github.com/4color/SiSaiBlog/api/blog"
	"github.com/4color/SiSaiBlog/api/checkcode"
	"github.com/4color/SiSaiBlog/api/login"
	"github.com/4color/SiSaiBlog/views"
	"github.com/4color/SiSaiBlog/views/adminview"
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
	R.LoadHTMLGlob("web/template/**/*")

	R.GET("/", views.Index)
	R.GET("/view/*blogurl", views.Viewblog)
	R.GET("/login", views.Login)
	R.GET("/logout", views.AdminLogout)

	store := cookie.NewStore([]byte("secret"))
	R.Use(sessions.Sessions("mysession", store))

	//管理员界面
	authorized := R.Group("/admin", api.AuthSessionMiddle())
	{

		authorized.GET("/", views.Admin)
		authorized.GET("/list", views.AdminListBlog)
		authorized.GET("/blogedit", views.AdminNew)
		authorized.GET("/reset", views.AdminReset)
		authorized.GET("/nav", adminview.List)
		authorized.GET("/navedit", adminview.Edit)

	}

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

		nav := v1.Group("/nav")
		{
			nav.GET("/list", apinav.List)
			nav.GET("/count", apinav.ListCount)
			nav.GET("/read/:id", apinav.Read)

			authorizednav := nav.Group("/", api.AuthSessionMiddle())
			{
				authorizednav.POST("/save", apinav.Save)
				authorizednav.POST("/delete", apinav.Delete)
			}
		}

		authorizedapi := v1.Group("/", api.AuthSessionMiddle())
		{
			authorizedapi.POST("/blogs/delete", blog.BlogDelete)
			authorizedapi.POST("/blog", blog.SaveBlog)
			authorizedapi.POST("/blogsetting", apisetting.SaveSettingInfo)
			authorizedapi.POST("/repwd", login.EditPwd)

		}

	}
}
