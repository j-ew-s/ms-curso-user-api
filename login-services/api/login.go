package api

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/j-ew-s/ms-curso-user-api/common"
	"github.com/j-ew-s/ms-curso-user-api/database"
	loginModels "github.com/j-ew-s/ms-curso-user-api/login-services/models"
	userModels "github.com/j-ew-s/ms-curso-user-api/user-services/models"
	"github.com/valyala/fasthttp"
)

type LoginServices struct {
	DBConn *database.SQLCommand
}

func (ls LoginServices) Login(ctx *fasthttp.RequestCtx) {

	login := loginModels.Login{}
	err := json.Unmarshal(ctx.PostBody(), &login)
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	result, err := ls.DBConn.ValidateLogin(login)
	if err != nil {
		fmt.Println(fmt.Printf(" Insert Error  :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	if result.Id > 0 {
		login, err := ls.RegisterLogin(result)
		if err != nil {
			fmt.Println(fmt.Printf(" Insert Error  :::  %+v ::", err))
			common.PrepareResponse(ctx, 500, nil)
		}
		if login.Id > 0 {
			common.PrepareResponse(ctx, 200, login.Token)
		} else {
			common.PrepareResponse(ctx, 500, nil)
		}
	}

	common.PrepareResponse(ctx, 401, nil)
}

func (ls LoginServices) RegisterLogin(user userModels.User) (*loginModels.RegisterLoging, error) {

	login := &loginModels.RegisterLoging{
		Username:   user.Username,
		UserId:     int32(user.Id),
		LogedAt:    time.Now(),
		LoginLimit: time.Now().Add(time.Hour * 1),
		Token:      fmt.Sprint(uuid.New()),
	}

	err := ls.DBConn.Login(login)
	if err != nil {
		fmt.Println(fmt.Printf(" Insert  Login Error  :::  %+v ::", err))
		return nil, err
	}

	return login, nil
}

func (ls LoginServices) Logout(ctx *fasthttp.RequestCtx) {
	logout := loginModels.Logout{}
	err := json.Unmarshal(ctx.PostBody(), &logout)
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	err = ls.DBConn.Logout(&logout)
	if err != nil {
		fmt.Println(fmt.Printf(" Insert Error  :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	common.PrepareResponse(ctx, 200, nil)
}
