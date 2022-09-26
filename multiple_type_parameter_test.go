package main

import (
	"fmt"
	"testing"
)

func MultipleParameter[T1 any, T2 any](param1 T1, param2 T2, param3 T1) {
	fmt.Println(param1, param2)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParameter[string, int]("oman", 24, "haiii")
}
