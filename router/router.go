package router

import (
	"github.com/gin-gonic/gin"
	."github.com/yangyouwei/router-api/apis"
)
func InitRouter() *gin.Engine {
	router := gin.Default()
	//IndexApi为一个Handler
	router.POST("/line", AddlineApi)
	router.GET("/lines", GetlinesApi)
	router.GET("/line/:id", GetlineApi)
	router.PUT("/line/:id", ModlineApi)
	router.DELETE("/line/:id", DellineApi)
	router.POST("/applay", ApplayApi)
	router.POST("/reboot", RebootApi)
	return router
}
