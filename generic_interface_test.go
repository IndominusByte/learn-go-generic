package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T) T
}

func ChangeValue[T any](method GetterSetter[T], value T) T {
	method.SetValue(value)
	return method.GetValue()
}

type MySetterGetter[T any] struct {
	Value T
}

func (d *MySetterGetter[T]) GetValue() T {
	return d.Value
}

func (d *MySetterGetter[T]) SetValue(value T) T {
	d.Value = value
	return d.Value
}

func TestGetterSetter(t *testing.T) {
	data := MySetterGetter[string]{Value: "oman"}

	assert.Equal(t, "pradipta", ChangeValue[string](&data, "pradipta"))
}
