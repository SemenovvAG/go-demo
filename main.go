package main

import (
	"fmt"
	"math"
)

func main() {

	var userHeight = 1.83
	var userKg float64 = 81
	fmt.Print("__ Калькулятор индекса массы тела __\n")
	fmt.Print("Введите свой рост: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес: ")
	fmt.Scan(userKg)
	var IMT = userKg / math.Pow(userHeight, 2)
	fmt.Print(IMT)
}
