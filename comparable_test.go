package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T1 comparable](value1, value2 T1) bool {
	return value1 == value2
}

func TestIsSame(t *testing.T) {
	assert.Equal(t, true, IsSame[string]("oman", "oman"))
	assert.Equal(t, false, IsSame[int](1, 2))
}
