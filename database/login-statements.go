package database

import (
	"fmt"

	loginModels "github.com/j-ew-s/ms-curso-user-api/login-services/models"
	userModel "github.com/j-ew-s/ms-curso-user-api/user-services/models"
)

func (sqlCommand SQLCommand) ValidateLogin(login loginModels.Login) (userModel.User, error) {

	user := &userModel.User{}
	db, err := sqlCommand.ExecuteSQLCommand()
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		return *user, err
	}

	err = db.Table("users").
		Select("id, username").
		Where("username = ? AND password =?", login.Username, login.Password).
		Scan(&user).
		Error

	return *user, err
}

func (sqlCommand SQLCommand) Login(login *loginModels.RegisterLoging) error {

	db, err := sqlCommand.ExecuteSQLCommand()
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		return err
	}

	err = db.Table("login").Create(&login).Error

	return err
}

func (sqlCommand SQLCommand) Logout(input *loginModels.Logout) error {

	db, err := sqlCommand.ExecuteSQLCommand()
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		return err
	}
	logout := loginModels.Logout{}
	db.Table("login").Delete(&logout, input)

	return nil
}
