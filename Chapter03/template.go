package main

import (
	"fmt"
)



// The 'abstract' MasterAlgorithm
type MasterAlgorithm struct {
	template Template
}

func (c *MasterAlgorithm) TemplateMethod() {
	// orchestrate the steps
	c.template.Step1()
	c.template.Step2()
}


// The steps which can be specialized
type Template interface {
	Step1()
	Step2()
}


// Variant A
type VariantA struct{}
func (c *VariantA) Step1() {
	fmt.Println("VariantA step 1")
}
func (c *VariantA) Step2() {
	fmt.Println("VariantA step 2")
}

func main() {
	masterAlgorithm := MasterAlgorithm{new(VariantA)}
	masterAlgorithm.TemplateMethod()
}