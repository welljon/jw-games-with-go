package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b *badGuy) { // No need to dereference the variables with a * in a struc as go does this for you
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func addOne(x int) {
	x = x + 1
}

func addOnePtr(x *int) {
	*x = *x + 1
}

func main() {
	x := 5
	fmt.Println(x)
	// &variable tells go to get the address in memory, where the variable lives. eg below is where x lives.
	xptr := &x
	// Can also be written var xptr *int = &x The *int lets go know that you want the pointer
	fmt.Println(xptr)

	addOne(x) //This does nothing to the x further up the function as the x it sends to addOne is a copy not the actual x
	fmt.Println(x)

	addOnePtr(xptr) // This one uses the pointer to reference the x rather than the copy
	fmt.Println(x)  // Pointers allow you to change the variable you pass to the function and not the copy of it

	p := position{4, 2}
	//fmt.Println(p.x)

	b := badGuy{"jojo", 100, p}
	whereIsBadGuy(&b) // The & lets it know that it's a pointer to the variable b
}
