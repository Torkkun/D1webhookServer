package server

import (
	"app/pkg/domain"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Engine *gin.Engine
}

func NewRouter() *Router {
	router := &Router{
		Engine: gin.Default(),
	}
	router.initRoute()
	return router
}

func (router *Router) Run() {
	router.Engine.Run()
}

func (router *Router) initRoute() {
	router.Engine.POST("/googleAssitant", func(ctx *gin.Context) {
		googleWebhookHandler(ctx)
	})

	router.Engine.POST("/meraki", func(ctx *gin.Context) {
		merakiWebhookHandler(ctx)
	})

	router.Engine.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(200, domain.TestResponse{Message: "get test ok"})
	})
	router.Engine.POST("/test", func(ctx *gin.Context) {
		test := domain.TestRequest{}
		if err := ctx.Bind(&test); err != nil {
			log.Println(err)
			ctx.JSON(400, "invalid response")
			return
		}
		if test.Test == "" {
			ctx.JSON(400, "not string")
			return
		}
		fmt.Println(test.Test)
		ctx.JSON(200, domain.TestResponse{Message: "post test ok"})
	})
}
