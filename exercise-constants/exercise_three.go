package main

import (
	"fmt"
	"math"
)

const pie float64 = math.Pi
const name = "no name"
const (
	const_1 = iota
	const_2
	const_3 string = string(iota)

)

func main() {

	fmt.Println(pie)
	fmt.Println(name)
	fmt.Println(const_1, const_2, const_3)
}
