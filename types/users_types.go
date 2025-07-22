package types

import (
	"pkg.formatio/lib"
	"pkg.formatio/prisma/db"
)

type ListUsersArgs struct {
	lib.BaseListFilterArgs
}

type CreateUserArgs struct {
	FirstName string
	LastName  string
	Email     string
	Password  string
	Roles     []string
}

type GetUserArgs struct {
	ID    *string
	Email *string
}

type UpdateUserArgs struct {
	ID        string    `swaggerignore:"true"`
	FirstName *string   `json:"firstName" swag-validate:"optional"`
	LastName  *string   `json:"lastName" swag-validate:"optional"`
	Email     *string   `json:"email" validate:"email,optional"`
	Password  *string   `json:"password" swag-validate:"optional"`
	Roles     *[]string `json:"roles" swaggerignore:"true"`
}

type DeleteUserArgs struct {
	ID string
}

type RegisterUserArgs struct {
	FirstName string `json:"firstName" swag-validate:"optional"`
	LastName  string `json:"lastName" swag-validate:"optional"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type RegisterUserResult struct {
	db.UserModel
	Tokens lib.AuthTokens `json:"tokens"`
}

type LoginUserArgs struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserResult struct {
	Tokens lib.AuthTokens `json:"tokens"`
}

type Auth0UserArgs struct {
	Token string `json:"token" validate:"required" `
}

type RefreshAccessTokenArgs struct {
	RefreshToken string `validate:"required" json:"refreshToken"`
}
