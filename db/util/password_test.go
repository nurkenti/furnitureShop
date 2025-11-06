package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPw(t *testing.T) {
	password := RandomString(6)

	HashPw, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, HashPw)

	err = CheckPassword(password, HashPw)
	require.NoError(t, err)

	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, HashPw)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	HashPw2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, HashPw)
	require.NotEqual(t, HashPw, HashPw2)
}
