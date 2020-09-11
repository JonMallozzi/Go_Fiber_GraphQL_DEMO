package postgres

import (
	"time"

	"github.com/JonMallozzi/Go_Fiber_GraphQL_DEMO/graph/model"
	"github.com/google/uuid"
)

func GraphQLInputUsertoDBUSer(newUser model.NewUser) *DbUser {
	return &DbUser{
		ID:          uuid.New(),
		Username:    newUser.Username,
		Password:    newUser.Password,
		Email:       newUser.Email,
		DateOfBirth: newUser.DateOfBirth,
		DateCreated: time.Now(),
	}

}

func DBUserToGraphQLUser(dbUser *DbUser) *model.User {
	return &model.User{
		ID:          dbUser.ID.String(),
		Username:    dbUser.Username,
		Password:    dbUser.Password,
		Email:       dbUser.Email,
		DateOfBirth: dbUser.DateOfBirth,
		DateCreated: dbUser.DateCreated,
	}
}
