package util

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
	return rand.Intn(99)
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

func RandomModel(a, b string) string {
	ab := a
	ba := b
	if rand.Intn(2) == 0 {
		return ab
	} else {
		return ba
	}
}
func RandomMaterial(a, b, c string) string {
	a1 := a
	b1 := b
	c1 := c
	r := rand.Intn(3)
	switch r {
	case 1:
		return a1
	case 2:
		return b1
	case 0:
		return c1
	}
	return ""
}
