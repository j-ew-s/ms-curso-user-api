package main

import (
	"log"

	"github.com/buaazp/fasthttprouter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/j-ew-s/ms-curso-user-api/database"
	userServices "github.com/j-ew-s/ms-curso-user-api/user-services"
	"github.com/valyala/fasthttp"
)

func main() {

	router := fasthttprouter.New()

	sqlCommand := SetDataBase()

	client := userServices.CatalogServiceMain(sqlCommand)
	userServices.SetRoutes(router, client)

	fasthttp.ListenAndServe(":5100", CORS(router.Handler))
}

var (
	corsAllowHeaders     = "*"
	corsAllowMethods     = "HEAD,GET,POST,PUT,DELETE,OPTIONS"
	corsAllowOrigin      = "*"
	corsAllowCredentials = "true"
)

// CORS handles CORS
func CORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Credentials", corsAllowCredentials)
		ctx.Response.Header.Set("Access-Control-Allow-Headers", corsAllowHeaders)
		ctx.Response.Header.Set("Access-Control-Allow-Methods", corsAllowMethods)
		ctx.Response.Header.Set("Access-Control-Allow-Origin", corsAllowOrigin)

		next(ctx)
	}
}

func SetDataBase() *database.SQLCommand {

	conn := database.SetDrivers("mysql", "user_management")

	sqlCommand := &database.SQLCommand{
		SqlConn: conn,
	}

	err := sqlCommand.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return sqlCommand

}
