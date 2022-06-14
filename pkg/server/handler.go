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
	ProjectorURL = "localhost:8000/projector/"
	MonitorURL   = ""
	SearchURL    = "localhost:8000/search/"
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
		// get
		GetProjector(ctx)
		ctx.JSON(200, domain.ResponsePayloadGoogleAssistant{
			Prompt: usecase.SuccessPrompt(),
		})
		return
	}
}

// GET
func GetProjector(ctx *gin.Context) {
	res, err := http.Get(ProjectorURL)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, domain.ResponsePayloadGoogleAssistant{
			Prompt: usecase.FailedPrompt("get出来ませんでした"),
		})
		return
	}
	defer res.Body.Close()
	var resmessage ResponseMessage
	if err := json.NewDecoder(res.Body).Decode(&resmessage); err != nil {
		log.Println(err)
		ctx.JSON(500, domain.ResponsePayloadGoogleAssistant{
			Prompt: usecase.FailedPrompt("サーバーエラー"),
		})
		return
	}
	log.Println(resmessage.Message)
}

// POST
func PostProjector(ctx *gin.Context) {
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
}

type ResponseMessage struct {
	Message string `json:"message"`
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
	// post
	//Postsearch(ctx, &reqp)

	// 一次的に
	// get
	res, err := http.Get(SearchURL)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	defer res.Body.Close()
	var resmessage ResponseMessage
	if err := json.NewDecoder(res.Body).Decode(&resmessage); err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	log.Println(resmessage.Message)
	ctx.JSON(200, "success")
}

// POST
func Postsearch(ctx *gin.Context, reqp *domain.RequestPayloadMeraki) {
	req := domain.DoorOpenRequest{}
	req.Time = reqp.OccuredAt
	jsonreq, err := json.Marshal(req)
	if err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	res, err := http.Post(SearchURL, "application/json", bytes.NewBuffer(jsonreq))
	if err != nil {
		log.Println(err)
		ctx.JSON(500, "Internal Server Error")
		return
	}
	defer res.Body.Close()
}
