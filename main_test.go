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

