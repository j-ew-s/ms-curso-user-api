package api

import (
	"encoding/json"
	"fmt"

	"github.com/j-ew-s/ms-curso-user-api/common"
	"github.com/j-ew-s/ms-curso-user-api/database"
	userModel "github.com/j-ew-s/ms-curso-user-api/user-services/models"
	"github.com/valyala/fasthttp"
)

type UserServices struct {
	DBConn *database.SQLCommand
}

func (us UserServices) Create(ctx *fasthttp.RequestCtx) {

	user := userModel.User{}
	err := json.Unmarshal(ctx.PostBody(), &user)
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	result, err := us.DBConn.InsertUser(user)
	if err != nil {
		fmt.Println(fmt.Printf(" Insert Error  :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	common.PrepareResponse(ctx, 201, result)
}

func (us UserServices) Update(ctx *fasthttp.RequestCtx) {
	user := userModel.User{}
	err := json.Unmarshal(ctx.PostBody(), &user)
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	result, err := us.DBConn.UpdatedUser(&user)
	if err != nil {
		fmt.Println(fmt.Printf(" Insert Error  :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	common.PrepareResponse(ctx, 200, result)
}

func (us UserServices) Delete(ctx *fasthttp.RequestCtx) {
	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	err := us.DBConn.DeleteUser(id)
	if err != nil {
		fmt.Println(fmt.Printf(" Insert Error  :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	common.PrepareResponse(ctx, 200, nil)
}

func (us UserServices) Detail(ctx *fasthttp.RequestCtx) {
	statusCode := 200

	id := fmt.Sprintf("%v", ctx.UserValue("id"))

	result, err := us.DBConn.GetUserByIDS(id)
	if err != nil {
		fmt.Println(fmt.Printf(" Insert Error  :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	if result.Id <= 0 {
		statusCode = 404
	}
	common.PrepareResponse(ctx, statusCode, result)
}

func (us UserServices) GetAll(ctx *fasthttp.RequestCtx) {

	statusCode := 200

	results, err := us.DBConn.GetAllUsers()
	if err != nil {
		fmt.Println(fmt.Printf(" Insert Error  :::  %+v ::", err))
		common.PrepareResponse(ctx, 500, nil)
	}

	if len(results) < 1 {
		statusCode = 404
	}

	common.PrepareResponse(ctx, statusCode, results)
}

func (us UserServices) Login(ctx *fasthttp.RequestCtx) {
	fmt.Println("Login")
}

func (us UserServices) Logout(ctx *fasthttp.RequestCtx) {
	fmt.Println("Logout")
}
