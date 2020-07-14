package document

import (
	"fmt"
	"math/rand"
	"time"
)

type Document interface {
	generate() Document
	valid() bool
}

type CPF struct {
	Identifier string
}

func (d *CPF) Generate() {
	base := seedBase()
	firstVerifier := calculateWeight(base)
	fmt.Println(firstVerifier)
}

func seedBase() []int {
	var base []int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 9; i++ {
		base = append(base, rand.Intn(9))
	}

	return base
}

func calculateWeight(identifer []int) int {
	var sum int

	for key, digit := range identifer {
		sum += digit * (10 - key)
	}

	return sum
}
