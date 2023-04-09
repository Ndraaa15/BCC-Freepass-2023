package handler

import (
	"bcc-freepass-2023/src/entity"
	"bcc-freepass-2023/src/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) StudenrLogin(ctx *gin.Context) {
	student := entity.Student{}
	studentInput := model.StudentLoginParam{}

	if err := ctx.ShouldBindJSON(&student); err != nil {

	}

	if err := ctx.ShouldBindJSON(&studentInput); err != nil {

	}
}
