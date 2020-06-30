package main

import "testing"

func TestStackOverflow(t *testing.T) {
	s:= New()
	for i:=0;i<=100;i++ {
		Err := s.Push(i)
		if Err != nil {
			if (Err != stackOverflow{}) {
			t.Errorf("Expected Stack to be Full , got %v",Err)
			}
		}
	}
}

func TestUnderOverflow(t *testing.T)  {
	s:= New()
    for i:=0;i<100;i++ {
		s.Push(i)
	}
	for i:=0;i<=100;i++ {
		_,Err := s.Pop()
		if Err != nil {
			if (Err != stackEmpty{}) {
				t.Errorf("Expected Stack to be empty, got %v",Err)
			} 
		} 

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
		Err := s.Push(tc.input)
		if Err != nil {
			t.Fatalf("Expected no error , but got %v",Err)
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
		val,Err := s.Pop()
        if val != output {
			t.Errorf("Expected %v , but got %v ",output,val)
		}
		if Err != nil {
			t.Errorf("Expected no error , but got %v ",Err)
		}
	}

}


func TestNormalPeep(t *testing.T) {
	s:= New()
	s.Push(5)
	output := 5
	val,Err := s.Peep()
     if val != output {
			t.Errorf("Expected %v , but got %v ",output,val)
		}
	if Err != nil {
			t.Errorf("Expected no error , but got %v ",Err)
	}	
}