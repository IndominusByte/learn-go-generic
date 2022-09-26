package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSample(t *testing.T) {
	result := Length[string]("Oman")
	assert.Equal(t, "Oman", result)

	resultNum := Length[int](100)
	assert.Equal(t, 100, resultNum)
}
