package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/ztjustin/questions_api/domain/users"
	"github.com/ztjustin/questions_api/services/user"
)

type UserHandler interface {
	GetAll(*gin.Context)
	FindById(*gin.Context)
	Create(*gin.Context)
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
		return
	}

	c.JSON(http.StatusOK, list)
}

func (handler *userHandler) FindById(c *gin.Context) {
	user, err := handler.service.FindById(c.Param("id_user"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (handler *userHandler) Create(c *gin.Context) {
	var user domain.User
	err := c.ShouldBindJSON(&user)

	if err != nil {
		c.JSON(http.StatusHTTPVersionNotSupported, err.Error())
		return
	}

	newUser, errInsert := handler.service.Create(&user)

	if errInsert != nil {
		c.JSON(http.StatusCreated, errInsert.Error())
	}

	c.JSON(http.StatusCreated, newUser)

}
