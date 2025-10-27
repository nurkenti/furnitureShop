package db

import (
	"math/rand"
	"strconv"
	"strings"
)

const alphabet = "abcdefghiklmnopqrstuvwxyz"

func RandomInt(seed int64) int {
	var r = rand.New(rand.NewSource(seed))
	rID := r.Intn(1000)
	return rID
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomID() int {
	return RandomInt(200)
}
func RandomPassword() string {
	number := RandomInt(80)
	p := strconv.Itoa(number)
	symbals := RandomString(3)
	return p + symbals
}

func RandomEmail() string {
	return RandomString(5) + "@mail.ru"
}

func RandomName() string {
	return RandomString(6)
}

func RandomAge() int {
	return RandomInt(2)
}

func RandomRole() string {
	switch rand.Intn(2) {
	case 0:
		return "customer"
	case 1:
		return "admin"
	case 2:
		return "customer"
	}
	return ""
}
