package tests

// import (
// 	"context"
// 	"log"
// 	"testing"

// 	"github.com/samber/lo"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/testcontainers/testcontainers-go/modules/postgres"

// 	"pkg.formatio/dao"
// 	"pkg.formatio/tests/config"
// 	_ "pkg.formatio/tests/config/init"
// 	"pkg.formatio/types"
// )

// func userDAORunner(t *testing.T, callback interface{}) {
// 	ctx, container := config.GetSetup()

// 	err := container.Restore(ctx, postgres.WithSnapshotName("genesis"))
// 	if err != nil {
// 		log.Fatal("[userDAORunner]: ", err)
// 	}

// 	config.UseRunner(
// 		t,
// 		callback,
// 		config.NewEnv,
// 		config.NewTestDatabaseConnection,
// 		dao.NewUserDao,
// 		func() (context.Context, *postgres.PostgresContainer) {
// 			return ctx, container
// 		},
// 	)
// }

// func TestDAOListUsers(t *testing.T) {
// 	userDAORunner(
// 		t,
// 		func(userDao dao.IUserDao) {
// 			_, err := userDAO.CreateUser(types.CreateUserArgs{
// 				FirstName: "john",
// 				LastName:  "doe",
// 				Email:     "john@example.com",
// 				Password:  "password",
// 				Roles:     []string{},
// 			})
// 			assert.Nil(t, err)

// 			result, err := userDAO.ListUsers(types.ListUsersArgs{})
// 			assert.Nil(t, err)

// 			assert.Equal(t, 1, len(result))
// 		})
// }

// func TestDAOCreateUser(t *testing.T) {
// 	userDAORunner(
// 		t,
// 		func(userDao dao.IUserDao) {
// 			expected := types.CreateUserArgs{
// 				FirstName: "john",
// 				LastName:  "doe",
// 				Email:     "john@example.com",
// 				Password:  "password",
// 				Roles:     []string{},
// 			}
// 			result, err := userDAO.CreateUser(expected)
// 			assert.Nil(t, err)

// 			resultFirstName, _ := result.FirstName()
// 			resultLastName, _ := result.LastName()
// 			resultEmail, _ := result.Email()
// 			resultRoles := result.Roles
// 			resultPassword, _ := result.Password()

// 			assert.Equal(t, expected.FirstName, resultFirstName)
// 			assert.Equal(t, expected.LastName, resultLastName)
// 			assert.Equal(t, expected.Email, resultEmail)
// 			assert.Equal(t, expected.Roles, resultRoles)
// 			assert.Equal(t, expected.Password, resultPassword)
// 		})
// }

// func TestDAOGetUser(t *testing.T) {
// 	userDAORunner(
// 		t,
// 		func(userDao dao.IUserDao) {
// 			expected := types.CreateUserArgs{
// 				FirstName: "john",
// 				LastName:  "doe",
// 				Email:     "john@example.com",
// 				Password:  "password",
// 				Roles:     []string{},
// 			}

// 			userDAO.CreateUser(expected)
// 			result, err := userDAO.GetUser(types.GetUserArgs{Email: &expected.Email})
// 			if err != nil {
// 				t.Log(err)
// 			}

// 			resultFirstName, _ := result.FirstName()
// 			resultLastName, _ := result.LastName()
// 			resultEmail, _ := result.Email()
// 			resultRoles := result.Roles
// 			resultPassword, _ := result.Password()

// 			assert.Equal(t, expected.FirstName, resultFirstName)
// 			assert.Equal(t, expected.LastName, resultLastName)
// 			assert.Equal(t, expected.Email, resultEmail)
// 			assert.Equal(t, expected.Roles, resultRoles)
// 			assert.Equal(t, expected.Password, resultPassword)
// 		},
// 	)
// }

// func TestDAOUpdateUser(t *testing.T) {
// 	userDAORunner(
// 		t,
// 		func(userDao dao.IUserDao) {
// 			expected := types.CreateUserArgs{
// 				FirstName: "john",
// 				LastName:  "doe",
// 				Email:     "john@example.com",
// 				Password:  "password",
// 				Roles:     []string{},
// 			}
// 			user, err := userDAO.CreateUser(expected)
// 			if err != nil {
// 				t.Log(err)
// 			}

// 			expected.Email = "user@example.com"
// 			result, err := userDAO.UpdateUser(types.UpdateUserArgs{
// 				ID:    user.ID,
// 				Email: lo.ToPtr(expected.Email),
// 			})
// 			if err != nil {
// 				t.Log(err)
// 			}

// 			resultEmail, _ := result.Email()

// 			assert.Equal(t, expected.Email, resultEmail)
// 		},
// 	)
// }

// func TestDAODeleteUser(t *testing.T) {
// 	userDAORunner(
// 		t,
// 		func(userDao dao.IUserDao) {
// 			expected := types.CreateUserArgs{
// 				FirstName: "john",
// 				LastName:  "doe",
// 				Email:     "john@example.com",
// 				Password:  "password",
// 				Roles:     []string{},
// 			}
// 			user, err := userDAO.CreateUser(expected)
// 			assert.Nil(t, err)

// 			if user != nil {
// 				err = userDAO.DeleteUser(types.DeleteUserArgs{ID: user.ID})
// 				assert.Nil(t, err)

// 				email, _ := user.Email()
// 				result, err := userDAO.GetUser(types.GetUserArgs{Email: &email})
// 				assert.NotNil(t, err)

// 				assert.Nil(t, result)
// 			}
// 		},
// 	)
// }
