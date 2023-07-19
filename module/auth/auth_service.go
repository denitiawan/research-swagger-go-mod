package auth

import (
	"denitiawan/research-swagger-gomod-gin/module/user"
)

type AuthService interface {
	Login(requestDto LoginDto) user.UserDto
}
