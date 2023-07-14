package util

import (
	"math/rand"
	"os"
	"time"

	"github.com/oklog/ulid"
)

func SetDBURL() string {
	url := os.Getenv("url")
	token := os.Getenv("token")
	s := url + "?authToken=" + token
	return s

}

func GenerateULID() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	uid := ulid.MustNew(ms, entropy)
	return uid.String()
}
