package userServices

import (
	"github.com/buaazp/fasthttprouter"

	"github.com/j-ew-s/ms-curso-user-api/database"
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

func SetRoutes(router *fasthttprouter.Router, client *api.UserServices) {
	router.POST("/", client.Create)
	router.PUT("/:id", client.Update)
	router.DELETE("/:id", client.Delete)
	router.GET("/", client.GetAll)
	router.GET("/:id", client.Detail)
}
