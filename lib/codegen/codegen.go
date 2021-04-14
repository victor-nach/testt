package codegen

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}


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
		randIndex := rand.Intn(len(letters))
		s[i] = letters[randIndex]
	}

	return string(s)
}
