package main

import (
	"fmt"
)

type stack struct {
	st [100]int
	top int 
}

func New() *stack {
	return &stack{ top : -1}
}

func main() {
	s := stack{}
}
