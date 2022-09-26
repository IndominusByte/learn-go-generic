package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// type approximation
type Test int

type Num interface {
	~int | int16
}

type Number interface {
	int | int8 | int16 | int32 | int64 | float32
}

func Min[T Number](num1, num2 T) T {
	if num1 < num2 {
		return num1
	}
	return num2
}

func Max[T Num](num1, num2 T) T {
	if num1 > num2 {
		return num1
	}
	return num2
}

func TestMinTypeSet(t *testing.T) {
	assert.Equal(t, 10, Min[int](10, 20))
	assert.Equal(t, float32(1.1), Min[float32](10.0, 1.1))
	// error: Min[float64](10.9, 10.2)
	assert.Equal(t, Test(20), Max[Test](10, 20))
}

func TestMinTypeInference(t *testing.T) {
	assert.Equal(t, 10, Min(10, 20))
	assert.Equal(t, float32(1.1), Min(float32(10.0), float32(1.1)))
	// error: Min[float64](10.9, 10.2)
	assert.Equal(t, Test(20), Max(Test(10), Test(20)))
}
