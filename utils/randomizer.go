package utils

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// Generates random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// Generates an owner name
func RandomUsername() string {
	return RandomString(6)
}

// Generates a random money amount
func RandomMoney() int64 {
	return RandomInt(0, 10000)
}

// Generates a random currency
func RandomCurrency() string {
	currencies := []string{"USD", "IDR", "JPY"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

func Tambah(a, b int64) int64 {
	result := a + b
	fmt.Println(result)
	return result
}
