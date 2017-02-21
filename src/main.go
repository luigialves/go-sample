package main

import (
	"calculate"
)

func main() {
	calc := calculate.Values{ 
		Value1: 1, 
		Value2: 9, 
	}

	calc.SumValues()
}