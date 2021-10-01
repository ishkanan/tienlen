package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrdinal(t *testing.T) {
	assert.Equal(t, "1st", Ordinal(1))
	assert.Equal(t, "2nd", Ordinal(2))
	assert.Equal(t, "3rd", Ordinal(3))
	assert.Equal(t, "4th", Ordinal(4))
}

func TestPrettyPrint(t *testing.T) {
	type person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	me := person{Name: "Ish", Age: 35}
	assert.Equal(t, "{\n  \"name\": \"Ish\",\n  \"age\": 35\n}", PrettyPrint(me))
}
