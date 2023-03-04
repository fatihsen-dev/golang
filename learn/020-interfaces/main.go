package main

import "fmt"

type Shape interface {
	area() float64
	perimeter() float64
}

type Rectangle struct {
	width  float64
	height float64
}

// r parametresini Rectangle structundan türetiyoruz
func (r Rectangle) area() float64 {
	return r.width * r.height
}

// r parametresini Rectangle structundan türetiyoruz
func (r Rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func printShape(s Shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {
	r := Rectangle{width: 3, height: 4}
   printShape(r)
}