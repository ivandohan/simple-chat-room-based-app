package user

import "context"

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}

type Service interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginService(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
}

type CreateUserRequest struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserResponse struct {
	ID       string  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type LoginUserRequest struct {
	Email string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginUserResponse struct {
	accessToken string
	ID string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

