package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	FirstName T
	LastName  T
}

func (d *Data[_]) SayHello(param string) string {
	return "Hello " + param
}

func (d *Data[T]) ChangeFirstName(change T) T {
	d.FirstName = change
	return d.FirstName
}

func TestGenericStruct(t *testing.T) {
	data := Data[string]{
		FirstName: "oman",
		LastName:  "dewantara",
	}

	fmt.Println(data)
	assert.Equal(t, "Hello pradipta", data.SayHello("pradipta"))
	assert.Equal(t, "indah", data.ChangeFirstName("indah"))
}
