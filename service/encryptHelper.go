package service

import (
	"crypto/md5"
	"fmt"
)

func Str2MD5Str(input string) (output string) {
	btyeStr := []byte(input)
	b := md5.Sum(btyeStr)
	md5str := fmt.Sprintf("%x", b)

	return md5str
}
