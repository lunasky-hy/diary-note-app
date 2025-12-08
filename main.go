package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type question struct {
	Id int `json:"id"`
	QText string `json:"qtext"`
}

type daiary struct {
	Id int `json:"id"`
	Note string `json:"note"`
	UserId int `json:"userId"`
	Qustion question `json:"question"`
	CreatedAt time `json:"createdAt"`
}

var mockQuestions = []question{
	{Id: 1, QText: "今の気分は？"},
	{Id: 2, QText: "今日の夕食は？"},
	{Id: 3, QText: "〇〇でいい感じ！"},
}

var mockDiaries = []daiary{
	{Id: 1, Note: "最高!", UserId: 1, Question: { Id: 1, QText:  "今の気分は？" }, CreatedAt: date.new()}
}

func main() {
	// loggerとrecoveryミドルウェア付きGinルーター作成
	r := gin.Default()

	{
		v1 := r.Group("/v1")

		v1.GET("/api/questions", func(c *gin.Context) {
			c.JSON(http.StatusOK, mockQuestions)
		})
		v1.POST("/api/questions", func(c *gin.Context) {
			c.String(http.StatusAccepted, `sended`);
		})
		

		v1.GET("/api/diaries", func(c *gin.Context) {
			c.JSON(http.StatusOK, mockQuestions)
		})
		v1.POST("/api/diaries", func(c *gin.Context) {
			c.String(http.StatusAccepted, `sended`);
		})
	}

	// ポート8080でサーバー起動（デフォルト）
	// 0.0.0.0:8080（Windowsではlocalhost:8080）で待機
	err := r.Run()
	if err != nil {
		return
	}
}
