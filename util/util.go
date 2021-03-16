package util

import (
	"math/rand"
	"time"
)

func RandsString(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzASDFGHJKLQWERTYUIOPZXCVBNM")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
