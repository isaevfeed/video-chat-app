package hash

import (
	"math/rand"
	"time"
)

func HashGenerate(hashLen int32) string {
	rand.Seed(time.Now().Unix())
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]rune, hashLen)
	for i, _ := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
