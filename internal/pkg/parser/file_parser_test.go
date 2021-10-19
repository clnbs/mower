package parser_test

import (
	"jellysmack/internal/pkg/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtract(t *testing.T) {
	fileParser := parser.NewFileParser("../../../tests/test.txt")
	b, err := fileParser.Extract()
	assert.Nil(t, err, err)
	assert.NotNil(t, b)
}
