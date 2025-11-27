package main

import (
	"fmt"
	"math"
)

func main() {

	var userHeight = 1.83
	var userKg float64 = 81
	var IMT = userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}
