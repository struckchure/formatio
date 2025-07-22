package seeds

import (
	"log"

	"github.com/samber/lo"
	"pkg.formatio/dao"
	"pkg.formatio/lib"
	"pkg.formatio/services"
	"pkg.formatio/types"
)

// The function `UserSeed` registers a new user with default admin credentials and updates their roles
// to include the "ADMIN" role.
func UserSeed(
	env lib.Env,
	userService services.IUserService,
	userDao dao.IUserDao,
) {
	log.Println("planting seeds ... :)")

	user, err := userService.RegisterUser(types.RegisterUserArgs{
		Email:    env.DEFAULT_ADMIN_EMAIL,
		Password: env.DEFAULT_ADMIN_PASS,
	})
	if err != nil {
		log.Printf("could not plant seeds :( ... %s", err)
	}

	hasher := lib.NewHasher()
	if user != nil {
		log.Println("watering seeds ... :)")
		userDao.UpdateUser(types.UpdateUserArgs{
			ID:       user.ID,
			Roles:    lo.ToPtr([]string{"ADMIN"}),
			Password: lo.ToPtr(hasher.HashPassword(env.DEFAULT_ADMIN_PASS)),
		})

		log.Println("seeds have been planted ... ;)")
	}
}
