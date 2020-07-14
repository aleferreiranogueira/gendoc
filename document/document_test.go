package document

import (
	"testing"
)

func TestCpfGenerate(t *testing.T) {
	cpf := CPF{}

	cpf.generate()
	if cpf.Identifier != "123" {
		t.Errorf("CPF identifier invalid, got %v, wanted 123", cpf.Identifier)
	}
}
