package utils

import (
	"io"
	"crypto/rand"
	"encoding/base64"
	"crypto/md5"
	"encoding/hex"
)

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func GetUUID() string {
	bytes := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(bytes))
}
