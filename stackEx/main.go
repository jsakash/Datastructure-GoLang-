package main

import "fmt"

type Stack struct {
	Values []int
}

// PUSH (adding values)
func (s *Stack) push(v int) {
	s.Values = append(s.Values, v)
}

//POP (removing values)

func (s *Stack) pop() int {
	length := len(s.Values) - 1
	toPop := s.Values[length]
	s.Values = s.Values[:length]
	return toPop
}

func main() {

	myStack := Stack{}
	myStack.push(100)
	myStack.push(200)
	myStack.push(300)
	myStack.push(400)
	myStack.push(500)
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)
}
