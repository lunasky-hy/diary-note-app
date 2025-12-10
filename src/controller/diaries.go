package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lunasky-hy/dialy-note-app/src/model"
	"github.com/lunasky-hy/dialy-note-app/src/service"
)

type DiariesController struct {
	service service.DiaryService
}

func (qc DiariesController) Get(c *gin.Context) {
	ques, _ := qc.service.Find()
	c.JSON(http.StatusOK, ques)
}

func (qc DiariesController) Post(c *gin.Context) {
	var json model.Diary
	if err := c.ShouldBindJSON(&json); err != nil {
		return
	}
	qc.service.Create(json);
	c.String(http.StatusAccepted, `sended`);
}

func CreateDiaryController(service service.DiaryService) DiariesController {
	controller := DiariesController{service: service}
	return controller
}