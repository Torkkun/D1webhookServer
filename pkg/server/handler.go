package server

import (
	"app/pkg/domain"
	"app/pkg/usecase"
	"bytes"
	"encoding/json"
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
	// モニターハンドラーの処理
	case "monitor":
		log.Println("monitor")
		// post
		req := domain.MonitorRequest{}
		req.Notification = true
		jsonreq, err := json.Marshal(req)
		if err != nil {
			log.Println(err)
			ctx.JSON(500, domain.ResponsePayloadGoogleAssistant{
				Prompt: usecase.FailedPrompt("受信失敗"),
			})
			return
		}
		res, err := http.Post(ProjectorURL, "application/json", bytes.NewBuffer(jsonreq))
		if err != nil {
			log.Println(err)
			ctx.JSON(500, domain.ResponsePayloadGoogleAssistant{
				Prompt: usecase.FailedPrompt("post出来ませんでした"),
			})
			return
		}
		defer res.Body.Close()
		ctx.JSON(200, domain.ResponsePayloadGoogleAssistant{
			Prompt: usecase.SuccessPrompt(),
		})
		return
	// プロジェクターハンドラーの場合
	case "projector":
		log.Println("projector")
		// post
		req := domain.ProjectorRequest{}
		req.Notification = true
		jsonreq, err := json.Marshal(req)
		if err != nil {
			log.Println(err)
			ctx.JSON(500, domain.ResponsePayloadGoogleAssistant{
				Prompt: usecase.FailedPrompt("受信できませんでした"),
			})
			return
		}
		res, err := http.Post(ProjectorURL, "application/json", bytes.NewBuffer(jsonreq))
		if err != nil {
			log.Println(err)
			ctx.JSON(500, domain.ResponsePayloadGoogleAssistant{
				Prompt: usecase.FailedPrompt("送信できませんでした"),
			})
			return
		}
		defer res.Body.Close()
		ctx.JSON(200, domain.ResponsePayloadGoogleAssistant{
			Prompt: usecase.SuccessPrompt(),
		})
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
	req := domain.DoorOpenRequest{}
	req.Time = reqp.OccuredAt
	jsonreq, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	res, err := http.Post(DoorURL, "application/json", bytes.NewBuffer(jsonreq))
	if err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	defer res.Body.Close()
}
