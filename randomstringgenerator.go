package random

import (
	"crypto/rand"
	"encoding/hex"
	"math"
)

func NewRandomStringGenerator(length int) *RandomStringGenerator {
	return &RandomStringGenerator{length}
}

type RandomStringGenerator struct {
	length int
}

func (r *RandomStringGenerator) Generate() (string, error) {
	return generateString(r.length)
}

func GenerateString(length int) (string, error) {
	return generateString(length)
}

func generateString(length int) (string, error) {
	n := math.Floor(float64(length / 2))

	bytes := make([]byte, int(n))
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	return hex.EncodeToString(bytes), nil
}
