package main

type Stack struct {
	elements []string
}

type Move struct {
	moveElementsCount int
	fromStackNumber   int
	toStackNumber     int
}

func (s *Stack) head() string {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) pushBottom(element string) {
	s.elements = append([]string{element}, s.elements...)
}

func (s *Stack) pushTop(element string) {
	s.elements = append(s.elements, element)
}

func (s *Stack) pushTopAll(elements []string) {
	s.elements = append(s.elements, elements...)
}

func (s *Stack) pop(count int) []string {
	var popResult []string
	if len(s.elements) < count {
		popResult = s.elements
		s.elements = s.elements[:0]
	} else {
		popResult = s.elements[len(s.elements)-count:]
		s.elements = s.elements[0 : len(s.elements)-count]
	}

	return popResult
}
