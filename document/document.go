//Package document provides the document Types and a quick New for Documentation instances
package document

import (
	"errors"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// DocTypes Holds a map of string values for each Document type for easier usage.
// Use the method New(docType string) to get a fake document
type DocTypes map[string]reflect.Type

//New receives the document string and resolve it's instance based on the DocTypes map
func (register DocTypes) New(name string) (interface{}, error) {
	if typ, ok := register[name]; ok {
		return reflect.New(typ).Elem().Interface().(Document).Generate(), nil
	}

	return nil, errors.New("Could not instance doc with provided type")
}

//Set register a string to the Type of i
func (register DocTypes) Set(name string, i interface{}) {
	register[name] = reflect.TypeOf(i)
}

// Document provides methods for generating a fake document
type Document interface {
	Generate() Document
}

// BrazilContext refers to documents that are valid in Brazil
const BrazilContext = "brazil"

// Cpf provides a brazilian Fiscal Id document
type Cpf struct {
	ID      string
	Context string
}

//Generate creates a fake document
func (d *Cpf) Generate() Document {
	d.Context = BrazilContext
	d.ID = d.makeID(time.Now().UnixNano())
	return d
}

// Based on the CPF algorithm will generate a random valid CPF
func (d Cpf) makeID(seed int64) string {
	base := d.randomBase(seed)

	// First digit
	weight := d.calculateWeight(base)
	base = append(base, d.calculateVerifierDigit(weight))

	//Second Digit
	weight = d.calculateWeight(base)

	// Whole document
	identifier := append(base, d.calculateVerifierDigit(weight))

	var digits []string

	for _, digit := range identifier {
		digits = append(digits, strconv.Itoa(digit))
	}

	return strings.Join(digits, "")
}

// Provides a random 9 length integer to serve as the base for the document
func (d Cpf) randomBase(seed int64) []int {
	var base []int

	rand.Seed(seed)

	for i := 0; i < 9; i++ {
		base = append(base, rand.Intn(9))
	}

	return base
}

// Will sum each digit multiplied by the length (first digit has the 9 multipliyer, second has 10, etc)
// the sum represent the weight of the verifier digit (last 2 digits)
func (d Cpf) calculateWeight(identifer []int) int {
	var sum int

	for key, digit := range identifer {
		sum += digit * (len(identifer) + 1 - key)
	}

	return sum
}

// Based on the sum of the weight, returns the verifier digit
func (d Cpf) calculateVerifierDigit(weight int) int {
	rem := weight % 11

	if rem < 2 {
		return 0
	}

	return 11 - rem
}
