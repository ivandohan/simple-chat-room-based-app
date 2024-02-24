package user

import (
	"context"
	"golang-server/utils"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	AppSecretKey = "it's me, Dohan!"
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

type AppTokenClaims struct {
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (s *service) LoginService(ctxParam context.Context, request *LoginUserRequest) (*LoginUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctxParam, s.timeout)
	defer cancel()

	user, err := s.Repository.GetUserByEmail(ctx, request.Email)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	err = utils.CheckPassword(request.Password, user.Password)
	if err != nil {
		return &LoginUserResponse{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AppTokenClaims{
		ID: strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	signedString, err := token.SignedString([]byte(AppSecretKey))
	if err != nil {
		return &LoginUserResponse{}, err
	}

	return &LoginUserResponse{
		accessToken: signedString,
		Username: user.Username,
		ID: strconv.Itoa(int(user.ID)),
	}, nil
}