package user

import (
	"context"
	"golang-server/utils"
	"strconv"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctxParam context.Context, request *CreateUserRequest) (*CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctxParam, s.timeout)
	defer cancel()

	// TODO: Hashing password
	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	newUser := &User{
		Username: request.Username,
		Email: request.Email,
		Password: hashedPassword,
	}

	result, err := s.Repository.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	response := &CreateUserResponse{
		ID: strconv.Itoa(int(result.ID)),
		Username: result.Username,
		Email: result.Email,
	}

	return response, nil
}