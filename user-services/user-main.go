package userServices

import (
	"github.com/buaazp/fasthttprouter"
	authServices "github.com/j-ew-s/ms-curso-user-api/auth-services"
	"github.com/j-ew-s/ms-curso-user-api/database"
	loginApi "github.com/j-ew-s/ms-curso-user-api/login-services/api"
	"github.com/j-ew-s/ms-curso-user-api/user-services/api"
)

type Server struct {
	*api.UserServices
}

func UserServiceMain(conn *database.SQLCommand) *api.UserServices {
	userServices := &api.UserServices{
		DBConn: conn,
	}

	return userServices
}

func SetRoutes(router *fasthttprouter.Router, client *api.UserServices, loginClient *loginApi.LoginServices) {
	router.POST("/", authServices.AuthSessionValidator(client.Create, loginClient))
	router.PUT("/:id", authServices.AuthSessionValidator(client.Update, loginClient))
	router.DELETE("/:id", authServices.AuthSessionValidator(client.Delete, loginClient))
	router.GET("/", authServices.AuthSessionValidator(client.GetAll, loginClient))
	router.GET("/:id", authServices.AuthSessionValidator(client.Detail, loginClient))
}
