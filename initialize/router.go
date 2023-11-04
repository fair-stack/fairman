// Package initialize
// @program: fairman-gin
// @author: [lliuhuan](https://github.com/lliuhuan)
// @create: 2022-02-11 16:38
package initialize

import (
	"cnic/fairman-gin/middleware"
	"cnic/fairman-gin/router"
	"net/http"
	"time"

	_ "cnic/fairman-gin/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routers() *gin.Engine {
	var Router = gin.Default()

	// Cross domain
	Router.Use(middleware.Cors())
	// handlepanic
	Router.Use(middleware.GinRecovery(true))
	// Processing timeout
	//Router.Use(middleware.TimeoutMiddleware(time.Hour))
	// Processing logs
	Router.Use(middleware.DefaultLogger())
	// Token Bucket
	Router.Use(middleware.RateLimitMiddleware(time.Second))

	//systemV2Router := router.GroupApp.V2.System
	//softwareV2Router := router.GroupApp.V2.Software
	dockerV2Router := router.GroupApp.V2.Docker
	systemV2Router := router.GroupApp.V2.System

	Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "2.0.1",
		})
	})

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	RouterV1Group := Router.Group("v1")
	{
		PublicGroup := RouterV1Group.Group("")
		{
			PublicGroup.StaticFS("/upload", http.Dir("./upload"))

		}
		PrivateGroup := RouterV1Group.Group("")
		PrivateGroup.Use(middleware.JWTAuthMiddleware())
		{
		}
	}

	RouterV2Group := Router.Group("v2")
	{
		PublicGroup := RouterV2Group.Group("")
		{
			//systemV2Router.InitBaseRouter(PublicGroup) // System Common Method
			dockerV2Router.InitEndpointRouter(PublicGroup)
			systemV2Router.InitBaseRouter(PublicGroup)
			dockerV2Router.InitDockerRouter(PublicGroup)
			dockerV2Router.InitKubernetesRouter(PublicGroup)

		}
		PrivateGroup := RouterV2Group.Group("")
		PrivateGroup.Use(middleware.JWTAuthMiddleware())
		{
			dockerV2Router.InitSoftwareRouter(PrivateGroup) // Software management methods
			//softwareV2Router.InitTemplateRouter(PrivateGroup) // Template
			//systemV2Router.InitEndpointRouter(PrivateGroup)
		}
	}

	return Router
}
