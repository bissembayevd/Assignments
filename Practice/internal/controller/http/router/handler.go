package router

import (
	"Goprojects/Practice/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/html/atom"
)

type Handler struct {
	usecase.UserInterface
}

func NewHandler(v1 *gin.RouterGroup, useCase usecase.UserInterface) *Handler {
	handler := &Handler{useCase: useCase}
	v1.POST("/create", handler.CreateUser)
}
func (h *Handler) CreateUser(c *gin.Context) {
	type name struct {
		Name string `json:"name"`
	}
	if err := c.BindJSON(&input) err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := h.useCase.CreateUser(input.Name)
	c.JSON(http.StatusOK, gin.H{})
}
