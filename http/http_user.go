package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ztjustin/questions_api/services/user"
)

type UserHandler interface {
	GetAll(*gin.Context)
}

type userHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) UserHandler {
	return &userHandler{
		service: service,
	}
}

func (handler *userHandler) GetAll(c *gin.Context) {
	list, err := handler.service.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, list)
}
