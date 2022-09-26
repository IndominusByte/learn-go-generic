package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface{ int64 | float64 }](first, second T) T {
	if first < second {
		return first
	}
	return second
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, int64(20), FindMin[int64](int64(20), int64(30)))
	assert.Equal(t, float64(1.0), FindMin[float64](float64(1.0), float64(2.0)))
}

func GetFirst[T []E, E any](param T) E {
	return param[0]
}

func TestGetFirst(t *testing.T) {
	assert.Equal(t, "oman", GetFirst[[]string, string]([]string{"oman", "pradipta"}))
	assert.Equal(t, 1, GetFirst[[]int, int]([]int{1, 2, 3}))
}
