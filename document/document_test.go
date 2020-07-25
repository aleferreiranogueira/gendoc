package document

import (
	"fmt"
	"testing"
)

func TestShouldGenerateCpf(t *testing.T) {
	doc := new(Cpf)
	doc.Generate()
	fmt.Println(doc)
	if doc.Context != BrazilContext {
		t.Errorf("Invalid context for CPF")
	}

	if doc.ID == "" {
		t.Errorf("Invalid ID for CPF")
	}
}
func TestShouldMakeId(t *testing.T) {
	doc := Cpf{}
	var seed int64 = 10
	id := doc.makeID(seed)
	fmt.Println(id)

	if id != "27783173250" {
		t.Errorf("Document document number invalid, got %v, wanted 27783173250", len(id))
	}
}

func TestSeedBase(t *testing.T) {
	doc := new(Cpf)
	base := doc.randomBase(1)

	if len(base) != 9 {
		t.Errorf("Base document number invalid, got %v, wanted 9", len(base))
	}
}

func TestCalculateFirstWeight(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			210,
		},
		{
			[]int{1, 1, 1, 4, 4, 4, 7, 7, 7},
			162,
		},
		{
			[]int{7, 2, 5, 6, 1, 6, 8, 9, 2},
			269,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.input, tt.want)
		doc := new(Cpf)

		t.Run(testname, func(t *testing.T) {
			if result := doc.calculateWeight(tt.input); result != tt.want {
				t.Errorf("Total sum for identifier invalid, got %v, wanted %v", result, tt.want)
			}
		})
	}
}

func TestCalculateSecondWeight(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
			255,
		},
		{
			[]int{1, 1, 1, 4, 4, 4, 7, 7, 7, 3},
			204,
		},
		{
			[]int{7, 2, 5, 6, 1, 6, 8, 9, 2, 6},
			327,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.input, tt.want)
		doc := new(Cpf)

		t.Run(testname, func(t *testing.T) {
			if result := doc.calculateWeight(tt.input); result != tt.want {
				t.Errorf("Total sum for identifier invalid, got %v, wanted %v", result, tt.want)
			}
		})
	}
}

func TestCalculateFirstVerifierDigit(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{
			210,
			0,
		},
		{
			162,
			3,
		},
		{
			269,
			6,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.input, tt.want)
		doc := new(Cpf)

		t.Run(testname, func(t *testing.T) {
			if result := doc.calculateVerifierDigit(tt.input); result != tt.want {
				t.Errorf("Verifier digit invalid, got %v, wanted %v", result, tt.want)
			}
		})
	}
}
