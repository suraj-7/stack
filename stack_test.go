package main

import "testing"

func TestPush(t *testing.T) {
	tests := []struct {
		input  int
	}{
		{0},{1},{2},{3},
	}
    s:= New()
	for _, tc := range tests {
		r := s.push(tc.input)
		if r != nil {
			t.Fatalf("Expected no error , but got %v",r)
		}
	}
	
}

func TestPop(t *testing.T) {
	
		s:= New()
		_,err := s.pop()
		if err == nil {
			t.Errorf("Expected Stack is empty , but got %v",err)
		}
}

func TestPeep(t *testing.T) {
	    s:= New()
		_,err := s.peep()
		if err == nil {
			t.Errorf("Expected Stack is empty , but got %v",err)
		}
	
}