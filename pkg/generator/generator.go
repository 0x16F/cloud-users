package generator

import (
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
)

func NewString(length uint) string {
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	randomBytes := make([]byte, length)

	rand.Read(randomBytes)

	for i := uint(0); i < length; i++ {
		result[i] = alphabet[int(randomBytes[i])%len(alphabet)]
	}

	return string(result)
}

func NewHash(str ...string) string {
	hasher := sha512.New()

	for _, s := range str {
		hasher.Write([]byte(s))
	}

	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))
}
