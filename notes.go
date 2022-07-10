package main

import (
	"fmt"
	"time"
)

type Point struct {
	x int
	y int
}

// func (p Point) OffsetX(offset int) {
// 	p.x = p.x + offset
// }

func (p *Point) OffsetX(offset int) {
	p.x = p.x + offset
}
func main() {

	v := MyInt(3)
	fmt.Println(v)          // 3
	fmt.Println(v.Double()) // 6
	fmt.Println(v)          // 3

	p1 := Point{3, 3}
	fmt.Println(p1)

	p1.OffsetX(5)
	fmt.Println(p1)

	// p2 := &p1
	// p2.OffsetX(5)
	// fmt.Println(p1)

	go sayHello()
	go sayHello()
	go sayHello()

	time.Sleep(1000 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}

type MyInt int

func (mi MyInt) Double() int {
	return int(mi * 2)
}

/*
	Defining interface
*/

type Shape2D interface {
	// Area() float64
	// Perimeter() float64
	foo() bool
}

type Triangle struct {
}

func (t *Triangle) foo() bool { return false }

type Rectangle struct {
}

func (r *Rectangle) foo() bool { return false }

func (t *Triangle) Area() float64 {
	// blah blah
	return 0.0
}

func (t *Triangle) Perimeter() float64 {
	// blah blah
	return 0.0
}

// Type assertion for disambiguation
func DrawShape(s Shape2D) bool {
	rect, ok := s.(*Rectangle)
	if ok {
		DrawRect(rect)
	}
	tri, ok := s.(*Triangle)
	if ok {
		DrawTriangle(tri)
	}
	return false
}

func DrawRect(r *Rectangle)    {}
func DrawTriangle(t *Triangle) {}

/*
	Defining interface type
*/
type Speaker interface {
	Speak()
}
type Dog struct {
	name string
}

func (d *Dog) Speak() {
	fmt.Println(d.name)
}

func testInterfaceType() {
	var s1 Speaker
	d1 := Dog{"Shiba"}
	s1 = &d1
	s1.Speak()
}
