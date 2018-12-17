package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	fmt.Println("Asserting that 1 == 1")
	assert.Equal(t, 1, 1, "something has gone horribly wrong")
}

func TestReadParams(t *testing.T) {
	fmt.Println("Asserting that readParams returns a populated object")
	params := readParams()
	assert.NotEmpty(t, params, "readParams returned an empty object")
}
