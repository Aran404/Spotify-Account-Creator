package generator

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dchest/uniuri"
)

var RandomDomain = func() string {
	rand.Seed(time.Now().UnixNano())

	domains := [...]string{
		"@gmail.com",
		"@outlook.com",
		"@hotmail.com",
		"@yahoo.com",
	}

	domain := domains[rand.Intn(len(domains))]

	return domain
}

func GetDob() string {
	year := rand.Intn(20) + 1980
	month, day := rand.Intn(8)+1, rand.Intn(8)+1

	return fmt.Sprintf("%d-0%d-0%d", year, month, day)
}

func RandomEmail() string {
	domain := RandomDomain()
	prefix := uniuri.NewLen(8)

	return prefix + domain
}

func (in *Instance) CheckInstanceFields() {
	if in.DisplayName == EmptyString {
		in.SetDisplayName()
	}

	if in.Email == EmptyString {
		in.SetEmail()
	}

	if in.Password == EmptyString {
		in.SetPassword()
	}
}

func RandomPassword(n int) string {
	return uniuri.NewLen(n)
}

func RandomSuffix(n int) string {
	return uniuri.NewLen(n)
}

func RandomStringSuffix(n int) string {
	return uniuri.NewLen(n)
}

func RandomDisplayName(n int) string {
	return uniuri.NewLen(n)
}
