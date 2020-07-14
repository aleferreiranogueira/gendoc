package document

type Document interface {
	generate() Document
	valid() bool
}

type CPF struct {
	Identifier string
}

func (d *CPF) generate() {
	d.Identifier = "123"
}
