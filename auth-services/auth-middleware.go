package authServices

import (
	"github.com/j-ew-s/ms-curso-user-api/common"
	loginapi "github.com/j-ew-s/ms-curso-user-api/login-services/api"
	"github.com/valyala/fasthttp"
)

func AuthSessionValidator(requestHandler fasthttp.RequestHandler, client *loginapi.LoginServices) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		token := string(ctx.Request.Header.Peek("Authorization"))

		if token != "" {

			result, err := client.IsTokenValid(token)

			if err != nil {
				common.PrepareResponse(ctx, 500, "Error")
				return
			}

			if result {
				requestHandler(ctx)
			} else {
				common.PrepareResponse(ctx, 401, "Token não válido")
			}

		} else {
			common.PrepareResponse(ctx, 401, "Não tem acesso ao conteúdo")
		}
	}
}
