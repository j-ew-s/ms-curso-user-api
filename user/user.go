package user

import (
	"context"
)

type Server struct {
}

func (s *Server) GetUser(c context.Context, u UserId) (*User, error) {
	if u.Id == "1" {
		return &User{
			Id:    "1",
			Name:  "A",
			Email: "a@a.com",
			Address: &Address{
				Street:  "Rua A",
				ZipCode: "1234567",
			},
			Account: &Account{
				User:     "user",
				Password: "123",
			},
		}, nil
	}

	return nil, nil
}

func (s *Server) IsTokenValid(c context.Context, t Token) (*TokenValidation, error) {
	if t.Token == "123" {
		return &TokenValidation{
			Status:  true,
			Message: "Valid",
		}, nil
	}

	return &TokenValidation{
		Status:  false,
		Message: "Invalid",
	}, nil
}
