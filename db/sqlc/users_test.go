package db

import (
	"testing"

	"github.com/nurkenti/furnitureShop/util"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		ID:           util.RandomID(),
		Email:        util.RandomEmail(),
		FullName:     util.RandomName,
		PasswordHash: util.RandomPassword(),
	}
}
