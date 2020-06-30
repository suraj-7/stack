package main

import "testing"

func TestStackOverflow(t *testing.T) {
	s:= New()
	for i:=0;i<100;i++ {
		s.Push(i)
	}
	if tp := s.top;tp != 99 {
		t.Fatalf("Expected Stack to be Full,but got top as %v",tp)
	}
}

func TestUnderOverflow(t *testing.T) {
	s:= New()
    for i:=0;i<99;i++ {
		s.Push(i)
	}
	for i:=0;i<99;i++ {
		s.Pop()
	}
	if tp := s.top;tp != -1 {
		t.Fatalf("Expected Stack to be Empty, but got %v",s.data[tp])
	}
}

func TestPush(t *testing.T) {
	tests := []struct {
		input  int
	}{
		{0},{1},{2},{3},
	}
    s:= New()
	for _, tc := range tests {
		r := s.Push(tc.input)
		if r != nil {
			t.Fatalf("Expected no error , but got %v",r)
		}
	}
	
}


func TestNormalPop(t *testing.T) {
	s:= New()
	for i:=0 ;i<5;i++ {
		s.Push(i)
	}
	for i:=0 ;i<5;i++ {
		output := 4-i
		val,_ := s.Pop()
        if val != output {
			t.Errorf("Expected %v , but got %v ",output,val)
		}
	}

}


func TestNormalPeep(t *testing.T) {
	s:= New()
	s.Push(5)
	output := 5
	val,_ := s.Peep()
     if val != output {
			t.Errorf("Expected %v , but got %v ",output,val)
		}
}