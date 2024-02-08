package application

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMD5(source string) string {
	hash := md5.Sum([]byte(source))

	return hex.EncodeToString(hash[:])
}
