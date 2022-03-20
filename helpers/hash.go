package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}
