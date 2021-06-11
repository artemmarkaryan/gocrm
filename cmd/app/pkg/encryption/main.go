package encryption

import (
	"crypto/md5"
	"fmt"
)

func Encrypt(word, secret string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(word + secret)))
}

func Matches(word, secret, hash string) bool {
	return Encrypt(word, secret) == hash
}