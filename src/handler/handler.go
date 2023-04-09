package handler

import (
	"bcc-freepass-2023/src/usecase"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	http *gin.Engine
	log  *log.Logger
	uc   *usecase.Usecase
}

func Init(uc *usecase.Usecase, log *log.Logger) *Handler {
	return &Handler{
		http: gin.Default(),
		log:  log,
		uc:   uc,
	}
}

func (h *Handler) Routes() {
	//Testing
	h.http.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world",
		})
	})

}

func (h *Handler) Run() {
	h.http.Run(":8080")
}
