package modules

import (
	"biligo/constant"
	"biligo/log"
	"biligo/modules/app"
	"biligo/modules/system"
	"biligo/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 全局路由设置文件

// 使用 gin 创建一个新的路由
func NewRouter() *gin.Engine {
	r := gin.Default()
	return r
}

// 注册各个模块路由
func RegisterRouter(r *gin.Engine) *gin.Engine {
	log.Debug("注册404页面")
	r.NoRoute(Page404)

	log.Debug("注册首页")
	r.GET("/", Index)

	log.Debug("注册 system 模块 路由")
	sysGroup := r.Group("/api/system")
	system.RouteSys(sysGroup)

	log.Debug("注册 app 模块 路由")
	appGroup := r.Group("/api/app")
	app.RegisterApp(appGroup)

	return r
}

// 放个根的路由 防止 404
// （下面一行的写法没有任何用处，只是表明 Index 的用处）
// @router / [GET]
func Index(c *gin.Context) {
	util.SuccessResult("Hello BiliGo").ToJSON(c)
}

// 404 页面
func Page404(c *gin.Context) {
	log.Warn(fmt.Sprintf("404 Page not found - %s %s ", c.Request.URL, c.Request.UserAgent()))

	result := util.FailResultWithCodeAndMessage(constant.HttpPageNotFound,
		"page not found", nil)

	c.JSON(constant.HttpPageNotFound, result)
}