//all of the resolvers that UserRepo users
//to fetch the object from postgresql using graphql
package postgres

import (
	"fmt"
	"time"

	"github.com/JonMallozzi/Go_Fiber_GraphQL_DEMO/graph/model"
	"github.com/google/uuid"

	"github.com/go-pg/pg/v10"
)

type UserRepo struct {
	DB *pg.DB
}

func (u *UserRepo) GetUsers() ([]*model.User, error) {

	//orm selecting all users
	var usersCollection []*DbUser
	err := u.DB.Model(&usersCollection).Select()
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for i := 0; i < len(usersCollection); i++ {
		users = append(users, DBUserToGraphQLUser(usersCollection[i]))
	}

	return users, nil
}

func (u *UserRepo) GetUserByID(id string) (*model.User, error) {

	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	var dbUser *DbUser = &DbUser{ID: userId}
	err = u.DB.Model(dbUser).WherePK().Select()
	if err != nil {
		return nil, err
	}

	var graphUser *model.User
	graphUser = DBUserToGraphQLUser(dbUser)

	return graphUser, nil
}

func (u *UserRepo) GetUserByUsername(username string) (*model.User, error) {

	dbUser := new(DbUser)
	err := u.DB.Model(dbUser).Where("username = ?", username).Select()
	if err != nil {
		return nil, err
	}

	var graphUser *model.User
	graphUser = DBUserToGraphQLUser(dbUser)

	return graphUser, nil
}

func (u *UserRepo) CreateUser(newUser model.NewUser) (*model.User, error) {

	dbUser := GraphQLInputUsertoDBUSer(newUser)
	user := &model.User{
		ID:          dbUser.ID.String(),
		Username:    newUser.Username,
		Password:    newUser.Password,
		Email:       newUser.Email,
		DateOfBirth: newUser.DateOfBirth,
		DateCreated: dbUser.DateCreated,
	}

	_, err := u.DB.Model(dbUser).Insert()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) UpdateUser(id string, username *string, password *string, email *string, date_of_birth *time.Time) (*model.User, error) {

	userId, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	//fetching the user to be updated
	var dbUser *DbUser = &DbUser{ID: userId}
	err = u.DB.Model(dbUser).WherePK().Select()
	if err != nil {
		return nil, err
	}

	//add variables if they are not nil
	if username != nil {
		dbUser.Username = *username
	}
	if password != nil {
		dbUser.Password = *password
	}
	if email != nil {
		dbUser.Email = *email
	}
	if date_of_birth != nil {
		dbUser.DateOfBirth = *date_of_birth
	}

	user := DBUserToGraphQLUser(dbUser)

	_, err = u.DB.Model(dbUser).WherePK().UpdateNotZero()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserRepo) DeleteUser(id string) (string, error) {

	userId, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	var dbUser *DbUser = &DbUser{ID: userId}
	err = u.DB.Model(dbUser).WherePK().Select()
	if err != nil {
		return "", err
	}

	_, err = u.DB.Model(dbUser).WherePK().Delete()
	if err != nil {
		return "", err
	}

	returnString := fmt.Sprintf("User with the id:%s was successfully deleted", userId)

	return returnString, nil

}
