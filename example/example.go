package main

import (
	"fmt"
	"math"
)

func main() {
	example()
}

func example() {
	a := 0.4
	b := 0.8
	xn := 3.2
	xk := 6.1
	deltaX := 0.6
	x1 := 4.48
	x2 := 3.56
	x3 := 2.78
	x4 := 5.28
	x5 := 3.21
	x := xn
	for x <= xk {
		result := ((math.Pow(a, x)) + (math.Pow(b, x)))
		result = (result / (math.Log(a/b) / 10.0)) * math.Sqrt(a*b)
		fmt.Printf("x=%.2f, y=%.2f\n", x, result)
		x += deltaX
	}
	values := []float64{x1, x2, x3, x4, x5}

	for _, x := range values {
		result := ((math.Pow(a, x)) + (math.Pow(b, x)))
		result = (result / (math.Log(a/b) / 10.0)) * math.Sqrt(a*b)
		formattedResult := fmt.Sprintf("%.3f", result)
		fmt.Println("При x=", x, " y=", formattedResult)
	}

}
