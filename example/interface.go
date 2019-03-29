package main

import (
	"fmt"
)

type Shaper interface {
	Area() int
	Circumference() int
}

type Rect struct {
	Width int
	Height int
}

type Circle struct {
	Radius int
}

func (r *Rect) Area() int {
	return r.Width * r.Height
}

func (r *Rect) Circumference() int {
	return 2 * (r.Width + r.Height)
}

func (c *Circle) Area() int {
	return c.Radius * c.Radius
}

func (c *Circle) Circumference() int {
	return 2 * c.Radius
}

func showShaper(shaper Shaper){
	fmt.Println(shaper.Area())
	fmt.Println(shaper.Circumference())
}

func main() {
	r := Rect{10, 20}
	fmt.Printf("Rect w: %f, d: %f, Area: %f, Circumference: %f\n", r.Width, r.Height, r.Area(), r.Circumference())

	c := Circle{5}
	fmt.Printf("Circle r: %f, Area: %f, Circumference: %f\n", c.Radius, c.Area(), c.Circumference())

	showShaper(&r)
	showShaper(&c)
}
