package main

import "testing"

func (s *stack)TestPush(t *testing.T) {
	for i:=0;i<100;i++ {
		err := s.push(i)
		if err != nil {
			t.Errorf("Expected no error , but got %v",err)
			return
		}
	}
	
}

func (s *stack)TestPop(t *testing.T) {
	for i:=0;i<100;i++ {
		output:= 99-i
		val,err := s.pop()
		if err != nil {
			t.Errorf("Expected no error , but got %v",err)
			return
		}
	
		if val != output {
			t.Errorf("expected output to be %v , but got %v",output,val)
			return
		}
	}
	

}

func (s *stack)TestPeep(t *testing.T) {
		_,err := s.peep()
		if err != nil {
			t.Errorf("Expected no error , but got %v",err)
			return
		}
	
}