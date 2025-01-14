package pattern

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

// https://golangbyexample.com/visitor-design-pattern-go/

import "fmt"

type shape interface {
	getType() string
	accept(visitor)
}

type square struct {
	side int
}

func (s *square) accept(v visitor) {
	v.visitForSquare(s)
}

func (s *square) getType() string {
	return "Square"
}

type circle struct {
	radius int
}

func (c *circle) accept(v visitor) {
	v.visitForCircle(c)
}

func (c *circle) getType() string {
	return "Circle"
}

type rectangle struct {
	l int
	b int
}

func (t *rectangle) accept(v visitor) {
	v.visitForrectangle(t)
}

func (t *rectangle) getType() string {
	return "rectangle"
}

type visitor interface {
	visitForSquare(*square)
	visitForCircle(*circle)
	visitForrectangle(*rectangle)
}

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitForSquare(s *square) {
	//Calculate area for square. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for square")
}

func (a *areaCalculator) visitForCircle(s *circle) {
	//Calculate are for circle. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for circle")
}

func (a *areaCalculator) visitForrectangle(s *rectangle) {
	//Calculate are for rectangle. After calculating the area assign in to the area instance variable
	fmt.Println("Calculating area for rectangle")
}

type middleCoordinates struct {
	x int
	y int
}

func (a *middleCoordinates) visitForSquare(s *square) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *middleCoordinates) visitForCircle(c *circle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *middleCoordinates) visitForrectangle(t *rectangle) {
	//Calculate middle point coordinates for square. After calculating the area assign in to the x and y instance variable.
	fmt.Println("Calculating middle point coordinates for rectangle")
}

// func main() {
// 	square := &square{side: 2}
// 	circle := &circle{radius: 3}
// 	rectangle := &rectangle{l: 2, b: 3}
// 	areaCalculator := &areaCalculator{}
// 	square.accept(areaCalculator)
// 	circle.accept(areaCalculator)
// 	rectangle.accept(areaCalculator)

// 	fmt.Println()
// 	middleCoordinates := &middleCoordinates{}
// 	square.accept(middleCoordinates)
// 	circle.accept(middleCoordinates)
// 	rectangle.accept(middleCoordinates)
// }
