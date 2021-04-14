package codegen

import "math/rand"

type Codegen interface {
	Generate() string
}

type codeGen struct {
}

//New returns a new instance of codegen
func New() *codeGen {
	return &codeGen{}
}

func (c *codeGen) Generate() string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
	s := make([]rune, 6)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}
