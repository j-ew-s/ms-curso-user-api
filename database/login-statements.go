package database

import (
	"fmt"

	userModel "github.com/j-ew-s/ms-curso-user-api/user-services/models"
)

func (sqlCommand SQLCommand) LoginUserStatement(user userModel.User) string {
	return fmt.Sprintf(" INSERT INTO user_management.users(name, username, password, email) VALUES('%s', '%s', '%s', '%s')", user.Name, user.Username, user.Password, user.Email)
}
