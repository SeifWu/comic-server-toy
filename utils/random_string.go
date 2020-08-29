package util

import (
	"math/rand"
	"time"
)

func randomString(n int, source []byte) string {
	result := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = source[rand.Intn(len(source))]
	}

	return string(result)
}

// RandomString 随机字符串
func RandomString(n int) string {
	var letters = []byte("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM1234567890-!@#$%^&*.,")
	result := randomString(n, letters)

	return result
}

// RandomAuthCodeString 随机数字验证码
func RandomAuthCodeString(n int) string {
	var letters = []byte("1234567890")
	result := randomString(n, letters)

	return result
}
