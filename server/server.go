package server

import "go-chat/internal/domain"

type SocketService interface {
	GetOnlineUsers() []domain.User
}
