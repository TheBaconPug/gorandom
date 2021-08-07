package gorandom

import (
	"math/rand"
	"time"
	"unsafe"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

const numbers = "0123456789"

const (
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

func RandInt(c int) int {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, c)
	for i, cache, remain := c-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = numbers[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*int)(unsafe.Pointer(&b))
}

func RandStr(c int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, c)
	for i, cache, remain := c-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return *(*string)(unsafe.Pointer(&b))
}
