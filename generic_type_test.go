package main

import (
	"fmt"
	"testing"
)

type Custom interface {
	string | int
}

type Bag[T Custom] []T

func PrintBag[T Custom](param Bag[T]) {
	for _, value := range param {
		fmt.Println(value)
	}
}

func TestPrintBag(t *testing.T) {
	PrintBag(Bag[int]{1, 2, 34})
	PrintBag(Bag[string]{"oman", "pradipta", "dewantara"})
}
