package common

import (
	"math/rand"
	"time"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyandzABCDEFGHIJKLMNOPQRSTUVWXYANDZ")

func randSequence(n int) string {
	b := make([]rune, n)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for i := range b {
		b[i] = letter[r1.Intn(99999)%len(letter)]
	}
	return string(b)
}

func GenSalt(length int) string {
	if length < 50 {
		length = 50
	}
	return randSequence(length)
}
