package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(useRectanglerule(function, 0, 5, 10))
	fmt.Println(trapezoidRule(function, 0, 5, 10))
}

func function(x float64) float64 {
	return 1 + x + math.Sin(x*2)
}

type mathFunction func(float64) float64

func trapezoidRule(f mathFunction, xmin, xmax float64, numIntervalse int) float64 {

	dx := (xmax - xmin) / float64(numIntervalse)

	totalArea := .0
	x := xmin

	for i := 0; i < numIntervalse; i++ {
		totalArea = totalArea + dx*(f(x)+f(x+dx))/2
		x = x + dx
	}

	return totalArea
}

func useRectanglerule(f mathFunction, xmin, xmax float64, numIntervals int) float64 {

	// calculate reqtanle width
	dx := (xmax - xmin) / float64(numIntervals)

	// add reqtanlge area
	totalArea := 0.
	x := xmin

	for i := 0; i < numIntervals; i++ {
		totalArea = totalArea + dx*f(x)
		x = x + dx
	}

	return totalArea
}

func sliceArea(f mathFunction, x1, x2 float64, maxSliceError float64) float64 {
	return .0
}
