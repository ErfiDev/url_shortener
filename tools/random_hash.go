package tools

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
	"math/big"
)

func RandomHash(str string) (string, error) {
	hash := fmt.Sprintf("%x", sha1.Sum([]byte(str)))

	var finalHash string

	for i := 0; i < 6; i++ {
		bigI, err := rand.Int(rand.Reader, big.NewInt(int64(len(hash))))
		if err != nil {
			return "", err
		}

		randInt := bigI.Int64()
		if randInt != 0 {
			randInt -= 1
		}

		finalHash += string(hash[randInt])
	}

	return finalHash, nil
}
