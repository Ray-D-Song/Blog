package router

import (
	"github.com/thinkerou/favicon"
	"server/controller"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	G := gin.New()
	G.Use(favicon.New("./favicon.ico"))
	R := G.Group("/api/v1")
	R.GET("/rss", controller.GetRss)

	R.GET("/blog/list", controller.GetBlogList)
	R.GET("/blog/content", controller.GetBlog)
	R.PUT("/blog/hit", controller.AddHit)

	R.GET("/remark/list", controller.GetRemarkList)
	R.POST("/remark/add", controller.AddRemark)
	R.PUT("/remark/del", controller.DelRemark)
	return G
}
