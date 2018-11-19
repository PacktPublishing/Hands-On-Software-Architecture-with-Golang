package main

import (
	"fmt"
	"math"
	"sort"
)


type Strategy interface {
	FindBreadth([]int) int // the algorithm
}

// A O(nlgn) implementation
type NaiveAlgo struct {}
func (n *NaiveAlgo) FindBreadth(set []int) int {
	sort.Ints(set) 
	return  set[len(set)-1] - set[0]
}

// A O(n) implementation
type FastAlgo struct {}
func (n *FastAlgo) FindBreadth(set []int) int {
	min := math.MaxInt32
	max := math.MinInt64

	for _,x := range(set) {
		if x < min {
			min = x
		}
		if x > max {
			max =x
		}
	}

	return max - min
}


// The client is ignorant to the algorithm
func client(s Strategy) int {
	a := []int { -1, 10, 3, 1}
	return s.FindBreadth(a)
}



func main() {
	fmt.Println(client(&NaiveAlgo{}))
	fmt.Println(client(&FastAlgo{}))
}