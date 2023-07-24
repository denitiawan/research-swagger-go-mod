package auth

import (
	"denitiawan/research-swagger-gomod-gin/common/dto"
)

type AuthService interface {
	Login(requestDto LoginDto) *dto.ImplResponse
}
