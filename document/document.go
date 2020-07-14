package document

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type Document interface {
	generate() Document
	valid() bool
}

type CPF struct{}

func (d CPF) Generate() string {
	base := seedBase()
	// First digit
	weight := calculateWeight(base)
	base = append(base, calculateVerifierDigit(weight))

	//Second Digit
	weight = calculateWeight(base)

	// Whole document
	identifier := append(base, calculateVerifierDigit(weight))

	var digits []string

	for _, digit := range identifier {
		digits = append(digits, strconv.Itoa(digit))
	}

	return strings.Join(digits, "")
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
		sum += digit * (len(identifer) + 1 - key)
	}

	return sum
}

func calculateVerifierDigit(weight int) int {
	rem := weight % 11

	if rem < 2 {
		return 0
	}

	return 11 - rem
}
