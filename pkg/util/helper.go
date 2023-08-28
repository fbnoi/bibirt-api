package util

import (
	"crypto/rand"
	"encoding/base64"
)

func GetRandomStr(len int) string {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return base64.StdEncoding.EncodeToString(b)
}
