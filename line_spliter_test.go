package linespliter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type T1 struct {
	Name string `line:"0:10"`
}

func TestLineSpliter(t *testing.T) {
	t1 := T1{}
	err := Unmarshal("0123456789abcdefghijklnmopqrstuvwxyz", &t1)
	assert.NoError(t, err, "Error on unmarshal")
	v, err := Marshal(&t1)
	t.Error(fmt.Sprintf("%s\n%s\n", t1.Name, v))
}
