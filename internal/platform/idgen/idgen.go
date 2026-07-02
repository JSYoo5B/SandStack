package idgen

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type Generator interface {
	Hex(size int) string
}

type RandomGenerator struct{}

func Random() Generator {
	return RandomGenerator{}
}

func (RandomGenerator) Hex(size int) string {
	return RandomHex(size)
}

type FixedGenerator struct {
	Value string
}

func Fixed(value string) Generator {
	return FixedGenerator{Value: value}
}

func (g FixedGenerator) Hex(_ int) string {
	return g.Value
}

func RandomHex(size int) string {
	bytes := make([]byte, size)
	if _, err := rand.Read(bytes); err != nil {
		return time.Now().UTC().Format("20060102150405.000000000")
	}
	return hex.EncodeToString(bytes)
}
