// Плюсы:
// 1. Упрощает добавление операций, работающих со сложными структурами объектов.
// 2. Объединяет родственные операции в одном объекте.
// Минусы:
// 1. Если структура объектов изменяется часто, то каждый раз придётся обновлять всех посетителей.
// 2. Может нарушать инкапсуляцию, посетители могут требовать доступ к внутренним полям объектов, что нарушает их приватность.
// Применимость:
// 1. Когда нужно выполнить какую-то операцию над элементами сложной структуры объектов, не изменяя классы этих объектов.
// 2. Когда новые операции часто добавляются к объектам, и вы хотите избежать постоянного изменения их классов.
// Примеры:
// 1. В системах работы с документами (например, XML-документы) паттерн «Посетитель» может использоваться для обхода элементов
// и выполнения операций над ними (например, проверка, рендеринг, экспорт в разные форматы).

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Accept(visitor ShapeVisitor)
}

type Circle struct {
	Radius float64
}

func (c *Circle) Accept(visitor ShapeVisitor) {
	visitor.VisitCircle(c)
}

type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) Accept(visitor ShapeVisitor) {
	visitor.VisitRectangle(r)
}

type ShapeVisitor interface {
	VisitCircle(circle *Circle)
	VisitRectangle(rectangle *Rectangle)
}

type AreaCalculator struct {
	Area float64
}

func (a *AreaCalculator) VisitCircle(circle *Circle) {
	a.Area = math.Pi * circle.Radius * circle.Radius
	fmt.Printf("Area of circle: %.2f\n", a.Area)
}

func (a *AreaCalculator) VisitRectangle(rectangle *Rectangle) {
	a.Area = rectangle.Width * rectangle.Height
	fmt.Printf("Area of rectangle: %.2f\n", a.Area)
}

type PerimeterCalculator struct {
	Perimeter float64
}

func (p *PerimeterCalculator) VisitCircle(circle *Circle) {
	p.Perimeter = 2 * math.Pi * circle.Radius
	fmt.Printf("Perimeter of circle: %.2f\n", p.Perimeter)
}

func (p *PerimeterCalculator) VisitRectangle(rectangle *Rectangle) {
	p.Perimeter = 2 * (rectangle.Width + rectangle.Height)
	fmt.Printf("Perimeter of rectangle: %.2f\n", p.Perimeter)
}

func main() {
	circle := &Circle{Radius: 5}
	rectangle := &Rectangle{Width: 4, Height: 3}

	areaCalc := &AreaCalculator{}

	circle.Accept(areaCalc)
	rectangle.Accept(areaCalc)

	perimeterCalc := &PerimeterCalculator{}

	circle.Accept(perimeterCalc)
	rectangle.Accept(perimeterCalc)
}
