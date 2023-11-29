package tests

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBytes_Field(t *testing.T) {
	assert.Equal(t, [][]byte{[]byte("go"), []byte("java"), []byte("python")}, bytes.Fields([]byte("go java python")))
	assert.Equal(t, [][]byte{[]byte("go"), []byte("java"), []byte("python")}, bytes.Fields([]byte("\tgo  \f \u0085 \u00a0 java \n\rpython")))
	assert.Equal(t, [][]byte{}, bytes.Fields([]byte("\t \n\r    ")))
}
