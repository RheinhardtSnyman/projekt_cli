package main

type myTypeStack struct {
	stackSlice []myType
}

func (s *myTypeStack) Push(in myType) {
	s.stackSlice = append(s.stackSlice, in)
}

func (s *myTypeStack) Pop() myType {
	if len(s.stackSlice) == 0 {
		var empty myType
		return empty
	}
	out := s.stackSlice[len(s.stackSlice)-1]
	s.stackSlice = s.stackSlice[:len(s.stackSlice)-1]
	return out
}
