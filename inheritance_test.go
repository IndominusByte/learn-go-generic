package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
type Construct interface {
	GetName() string
}

type Hai1 struct {
	Name string
}

func (h *Hai1) GetName() string {
	return h.Name
}

type Hai2 struct {
	Name string
}

func (h *Hai2) GetName() string {
	return h.Name
}

func (h *Hai2) GetFullName() string {
	return "Hello " + h.Name
}

func Lol(param Construct) string {
	return param.GetName()
}

func TestInheritance(t *testing.T) {
	fmt.Println(Lol(&Hai1{Name: "oman"}))
	fmt.Println(Lol(&Hai2{Name: "oman"}))
}
*/

type Employee interface {
	GetName() string
}

func GetName[T Employee](param T) string {
	return param.GetName()
}

type Manager interface {
	GetName() string
	GetManagerName() string
}

type President interface {
	GetName() string
	GetPresidentName() string
}

type MyManager struct {
	Name string
}

func (m *MyManager) GetName() string {
	return m.Name
}

func (m *MyManager) GetManagerName() string {
	return m.Name
}

type MyPresident struct {
	Name string
}

func (m *MyPresident) GetName() string {
	return m.Name
}

func (m *MyPresident) GetPresidentName() string {
	return m.Name
}

func TestInheritance(t *testing.T) {
	assert.Equal(t, "oman", GetName[Manager](&MyManager{Name: "oman"}))
	assert.Equal(t, "pradipta", GetName[President](&MyPresident{Name: "pradipta"}))
}
