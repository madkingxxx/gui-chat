package gui

import (
	"go-chat/internal/domain"
)

func getUserList() []domain.User {
	return []domain.User{
		{
			Login: "matt11",
			Name:  "Mattias",
		},
		{
			Login: "joj1",
			Name:  "John",
		},
		{
			Login: "kenni",
			Name:  "Kenny",
		},
	}
}
