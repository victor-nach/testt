package codegen

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCodeGen_Generate(t *testing.T) {
	codeGen := New()
	code1 := codeGen.Generate()
	assert.NotNil(t, code1)

	code2 := codeGen.Generate()
	assert.NotEqual(t, code1, code2)
}
