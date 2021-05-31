package application

import (
	"github.com/gin-gonic/gin"
	"github.com/ztjustin/questions_api/http"
	user_repo "github.com/ztjustin/questions_api/repository/user"
	user_service "github.com/ztjustin/questions_api/services/user"
)

var (
	router = gin.Default()
)

func StartApplication() {
	userHandler := http.NewUserHandler(user_service.NewService(user_repo.NewUserRepository()))

	router.GET("/users", userHandler.GetAll)
	router.GET("/users/:id_user", userHandler.FindById)
	router.POST("/users/create", userHandler.Create)
	router.Run()

}
