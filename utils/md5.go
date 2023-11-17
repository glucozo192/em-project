package utils

import (
	"crypto/md5"
	"fmt"
	"time"
)

func GetMD5HashWithRandom(s string) string {
	t := time.Now().UnixNano()
	return fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", s, t))))
}
