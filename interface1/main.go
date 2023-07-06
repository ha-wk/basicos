/*package main

import "fmt"

// Dog is a type representing a dog.
type Dog struct {
	Name string
}

// Cat is a type representing a cat.
type Cat struct {
	Name string
}

// Speaker is an interface that defines the behavior of speaking.
type Speaker interface {
	Speak() string
}

// Speak implements the Speak method for Dog.
func (d Dog) Speak() string {
	return "Woof!"
}

// Speak implements the Speak method for Cat.
func (c Cat) Speak() string {
	return "Meow!"
}

func main() {
	// Declare a variable of type Speaker and assign a Dog instance to it.
	var speaker Speaker = Dog{Name: "Buddy"}
	fmt.Println(speaker.Speak()) // Output: Woof!

	// Assign a Cat instance to the same variable.
	speaker = Cat{Name: "Whiskers"}
	fmt.Println(speaker.Speak()) // Output: Meow!
}

*/

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

// Rectangle is a struct
type Rectangle struct {
	Width  float64
	Height float64
}

// returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// struct representing a circle.
type Circle struct {
	Radius float64
}

// returns the area of the circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circ := Circle{Radius: 2.5}

	// Both Rectangle and Circle implement the Shape interface,
	// so we can use the Shape interface type to store instances of Rectangle and Circle(i.e the values can be initaialize here only)
	
	
	shapes := []Shape{rect, circ}

	for _, shape := range shapes {
		area := shape.Area()
		fmt.Printf("Area: %.2f\n", area)
	}
}
