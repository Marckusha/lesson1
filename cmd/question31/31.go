package main

import (
	"fmt"
	"math"
)

/*
Написать программу нахождения расстояния между 2 точками, которые
представление в виде структуры Point с инкапсулированными параметрами x,y
и конструктором
*/

type Point struct {
	x, y float64
}

func (point1 Point) distance(point2 Point) float64 {
	return math.Sqrt(math.Pow((point1.x-point2.x), 2) + math.Pow((point1.y-point2.y), 2))
}

func makePoint(x, y float64) Point {
	return Point{x, y}
}

func main() {
	point1 := makePoint(-2.3, 4)
	point2 := makePoint(8.5, 0.7)

	fmt.Println(point1.distance(point2))
}
