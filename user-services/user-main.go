package userservices

import (
	"github.com/buaazp/fasthttprouter"

	"github.com/j-ew-s/ms-curso-user-api/user-services/api"
)

func SetRoutes(router *fasthttprouter.Router) {
	router.POST("/", api.Create)
	router.PUT("/:id", api.Update)
	router.DELETE("/", api.Delete)
	router.GET("/", api.Detail)
	router.POST("/", api.Login)
	router.POST("/", api.Logout)
}
