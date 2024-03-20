package main

import "fmt"

type Shape struct {
	Name string
}

func (s Shape) Getname() string {
	return s.Name
}
func (s Shape) Area() float64 {
	return 0.0
}

type Circle struct {
	Shape
	Radius float64
}
type Rectangle struct {
	Shape
	Width  float64
	Height float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}
func main() {
	rectangle := Rectangle{Shape{"Треугольник"}, 10.0, 20.0}
	circle := Circle{Shape{"Круг"}, 5.0}
	fmt.Printf("%s площадью %.2f\n", rectangle.Getname(), rectangle.Area())
	fmt.Printf("%s площадью %.2f\n", circle.Getname(), circle.Area())

}
