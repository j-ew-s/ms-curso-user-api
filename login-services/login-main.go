package loginServices

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/j-ew-s/ms-curso-user-api/database"
	"github.com/j-ew-s/ms-curso-user-api/login-services/api"
)

type Server struct {
	*api.LoginServices
}

func LoginserviceMain(conn *database.SQLCommand) *api.LoginServices {
	userServices := &api.LoginServices{
		DBConn: conn,
	}

	return userServices
}

func SetRoutes(router *fasthttprouter.Router, client *api.LoginServices) {
	router.POST("/login", client.Login)
	router.POST("/logout", client.Logout)
}
