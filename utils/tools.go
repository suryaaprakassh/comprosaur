package utils

import (
	"math/rand"
	"os/exec"
)

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


func IsInstalled(name string) bool {
	_,err := exec.LookPath(name)
	return err != nil
}
