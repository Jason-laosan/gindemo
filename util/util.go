package util

import "math/rand"

func RandomString(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	result := make([]byte, n)
	for v := range letters {
		result[v] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}
