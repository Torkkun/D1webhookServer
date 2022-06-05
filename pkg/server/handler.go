package server

import (
	"app/pkg/domain"
	"app/pkg/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	ProjectorURL = ""
	MonitorURL   = ""
	DoorURL      = ""
)

func googleWebhookHandler(ctx *gin.Context) {
	reqp := domain.RequestPayloadGoogleAssistant{}
	if err := ctx.Bind(&reqp); err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	switch reqp.Handler.Name {
	case "monitor":
		log.Println("monitor")
		ctx.JSON(200, domain.ResponsePayloadGoogleAssistant{
			Prompt: usecase.MonitorPrompt(),
		})
		return

	case "projector":
		log.Println("projector")
	}
}

func merakiWebhookHandler(ctx *gin.Context) {
	// もろもろの認証やらすっ飛ばしドアアラートだけ受け取る
	reqp := domain.RequestPayloadMeraki{}
	if err := ctx.Bind(&reqp); err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	// 色々すっ飛ばして起きた時間だけ送る
	resp := domain.DoorOpenRequest{}
	resp.Time = reqp.OccuredAt
	res, err := http.Post(DoorURL, "application/json")
}
