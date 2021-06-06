package jwt

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func simple(plain string) string {
	hash := sha1.New()
	hash.Write([]byte(plain))
	return hex.EncodeToString(hash.Sum(nil))
}

func hs256(plain string, secretKey string) string {
	hash := hmac.New(sha256.New, []byte(secretKey))
	hash.Write([]byte(plain))
	return hex.EncodeToString(hash.Sum(nil))
}
