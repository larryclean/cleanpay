package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
)

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func Sha256(str,apiKey string) string {
	s := hmac.New(sha256.New,[]byte(apiKey))
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}
