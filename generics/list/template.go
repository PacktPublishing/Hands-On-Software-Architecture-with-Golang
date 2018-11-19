package list

import "github.com/cheekybits/genny/generic"

// The gogenerate command is used to create multiple instantiations from the generic type
// NOTE the replacement of tag "Element" - It's important that all elements below start with that
//go:generate genny -in=template.go -out=list-unit.go gen "Element=uint"

// Element is a generic type which will be contained in the list
type Element generic.Type

// ElementList is the generic FIFO queue
type ElementList struct {
	list []Element
}

// NewElementList creates a new list
func NewElementList() *ElementList {
	return &ElementList{list: []Element{}}
}

// Add inserts an element to the end of the list
func (l *ElementList) Add(v Element) {
	l.list = append(l.list, v)
}

// Get pops the element from the head of the list
func (l *ElementList) Get() Element {
	r := l.list[0]
	l.list = l.list[1:]
	return r
}
