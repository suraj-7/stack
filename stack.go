package main

import (
	"fmt"
	"errors"
	"sync"
)

type stacks struct {
	st [100]int
	top int 
	mux sync.Mutex
}

func (s stacks) isEmpty() error {
	if s.top == -1 {
		return errors.New("Stack is Empty")
	}
	return nil
}

func (s *stacks) overflow() error {
	if (*s).top >= 100 {
		return errors.New("Stack is Full")
	}
	return nil
}

func (s *stacks) underflow() error {
	if (*s).top <= -1 {
		return errors.New("Stack is Empty")
	}
	return nil
}

func (s *stacks) push(x int) (int,error) {
	err := (*s).overflow()
	if err != nil {
		return -1,err
	}
	(*s).mux.Lock()
	(*s).top++
	(*s).st[(*s).top] = x
	defer (*s).mux.Unlock()
	return x,nil
}


func (s *stacks) pop() (int,error){
	err := (*s).underflow()
	if err != nil {
		return -1,err
	}
	(*s).mux.Lock()
	x := (*s).st[s.top]
	(*s).top--
	defer (*s).mux.Unlock()
	return x,nil
}

func (s stacks) peek() (int,error){
	err := s.isEmpty()
	if err != nil {
		return -1,err
	}
	return s.st[s.top],nil
}

func main() {
	s := stacks{ top : -1}
	fmt.Println(s.push(10))
	fmt.Println(s.push(15))
	fmt.Println(s.push(20))
	fmt.Println(s.pop())
	fmt.Println(s.peek())
	fmt.Println(s.pop())
	fmt.Println(s.peek())
	fmt.Println(s.pop())
	fmt.Println(s.pop())
}
