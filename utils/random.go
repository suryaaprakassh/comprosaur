package utils

import "math/rand"

const CharacterPool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
			idx := rand.Intn(len(CharacterPool))
			b[i] = CharacterPool[idx]
			i++
	}
	return string(b)
}

