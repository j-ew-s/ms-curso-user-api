package database

import (
	"fmt"

	userModel "github.com/j-ew-s/ms-curso-user-api/user-services/models"
)

func (sqlCommand SQLCommand) InsertUser(user userModel.User) (userModel.User, error) {

	db, err := sqlCommand.ExecuteSQLCommand()
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		return user, err
	}

	db.Table("users").Create(&user)

	return user, nil
}

func (sqlCommand SQLCommand) UpdatedUser(user *userModel.User) (userModel.User, error) {

	db, err := sqlCommand.ExecuteSQLCommand()
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		return *user, nil
	}
	db.Table("users").Create(&user)

	return *user, nil

}

func (sqlCommand SQLCommand) DeleteUser(id string) error {
	db, err := sqlCommand.ExecuteSQLCommand()
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		return err
	}
	user := userModel.User{}
	db.Table("users").Delete(&user, id)

	return nil
}

func (sqlCommand SQLCommand) GetAllUsers() ([]*userModel.User, error) {
	users := make([]*userModel.User, 0)

	db, err := sqlCommand.ExecuteSQLCommand()
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		return users, err
	}
	db = db.Select("*").Table("users")
	db.Scan(&users)
	return users, nil
}

func (sqlCommand SQLCommand) GetUserByIDS(id string) (user userModel.User, err error) {
	db, err := sqlCommand.ExecuteSQLCommand()
	if err != nil {
		fmt.Println(fmt.Printf(" ERROR :::  %+v ::", err))
		return user, err
	}
	db = db.Select("*").Table("users").Where("id = ?", id)
	db.Scan(&user)

	return user, nil
}
