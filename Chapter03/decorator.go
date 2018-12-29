package main

import (
	"fmt"
	"time"
	"math"
)

type Function func(float64) float64

// This decorate profiles the execution time of function fn
func ProfileDecorator(fn Function) Function {
	return func(params float64) float64 {
		start := time.Now()
		result := fn(params)
		elapsed := time.Now().Sub(start)
		fmt.Println("Funtion completed with time : ", elapsed)
		return result
	}
}


func SquareRoot(n float64) float64 {
	return math.Sqrt(n)
}

func main() {

	decoratedSqaureRoot := ProfileDecorator(SquareRoot)
	fmt.Println(decoratedSqaureRoot(16))

}
