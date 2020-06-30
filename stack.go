package main

import (
	"fmt"
)

type stack struct {
	data [100]int
	top int 
}

func New() *stack {
	return &stack{ top : -1}
}

type stackOverflow struct {}

func (so stackOverflow) Error() string {
	return "Stack is full!"
}

type stackEmpty struct {}

func (se stackEmpty) Error() string {
	return "stack is empty!"
}

func (s *stack) Push(elem int) error {
   if s.top == 99 {
	   return stackOverflow{}
   }
   s.top++
   s.data[s.top] = elem
   return nil
}

func (s *stack) Pop() (int,error) {
	if s.top == -1 {
		return -1,stackEmpty{}
	}
	elem := s.data[s.top]
    s.top--
	return elem,nil
}

func (s *stack) Peep() (int,error) {
	if s.top == -1 {
		return -1,stackEmpty{}
	}
	return s.data[s.top],nil 
}

func main() {
	s := New()
	err := s.Push(5)
	if err != nil {
		fmt.Println("Error Occured : ",err)
		return
	}
	fmt.Println("5 is pushed")
	val,err := s.Pop() 
    if err != nil {
		fmt.Println("Error Occured : ",err)
		return
	}
	fmt.Println("Poped Element : ",val)
	val,err = s.Peep()
	if err != nil {
		fmt.Println("Error Occured : ",err)
		return
	}
	fmt.Println("Element at top : ",val)
	
}
