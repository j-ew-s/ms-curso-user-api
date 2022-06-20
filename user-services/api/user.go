package api

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Create(ctx *fasthttp.RequestCtx) {
	fmt.Println("Create")
}

func Update(ctx *fasthttp.RequestCtx) {
	fmt.Println("Update")
}

func Delete(ctx *fasthttp.RequestCtx) {
	fmt.Println("Delete")
}

func Detail(ctx *fasthttp.RequestCtx) {
	fmt.Println("Detail")
}

func Login(ctx *fasthttp.RequestCtx) {
	fmt.Println("Login")
}

func Logout(ctx *fasthttp.RequestCtx) {
	fmt.Println("Logout")
}
