package main

import "fmt"

type shape interface {
	area() float32
}

// STRUCT untuk persegi
type rectangle struct {
	length float32
}

// METHOD untuk area persegi
func (r rectangle) area() float32 {
	return r.length * r.length
}

// STRUCT untuk segitiga
type triangle struct {
	base   float32
	height float32
}

// METHOD untuk area segitiga
func (t triangle) area() float32 {
	return t.base * t.height / 2
}

func calculateArea(s shape) float32 {
	// lakukan pengecekan s shape
	if r, ok := s.(rectangle); ok { // r itu nilai dari struct rectangle
		r.length += 5   // ubah nilai length di r
		return r.area() // kembalikan nilai luas dari r yg baru
	}
	if t, ok := s.(triangle); ok { // t itu nilai dari struct triangle
		t.base += 10    // ubah nilai base di t
		return t.area() // kembalikan nilai luas dari base t yg baru
	}
	return s.area()
}

func main() {
	r := rectangle{
		length: 10,
	}

	t := triangle{
		base:   10,
		height: 10,
	}
	fmt.Println(calculateArea(r))
	fmt.Println(calculateArea(t))
}
