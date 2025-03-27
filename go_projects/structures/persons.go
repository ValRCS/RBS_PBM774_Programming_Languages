//let's explore making our own structures in Go
//we can create a structure in Go using the struct keyword
//a structure is a user defined type in Go
//for example, let's create a structure called person

package main

import (
	"fmt"
)

//a structure is a user defined type in Go
//for example, let's create a structure called person
//a person has a name and age
//we can create a structure in Go using the struct keyword

type Person struct {
	Name string
	Age  int
}

//let's create a structure method for Person called greet
//a structure method is a function that is associated with a structure
//a structure method is defined using the func keyword followed by the name of the method

func (p Person) greet() {
	fmt.Println("Hello, my name is", p.Name)
	//show age
	fmt.Println("I am", p.Age, "years old")
	fmt.Println("You killed my father, prepare to die") //this is Princess Bride reference of course
}

//let's add a birthday method to the Person structure which will modify the age of the person
func (p *Person) birthday() { //note the pointer receiver, with pointers we can modify the value of the structure
	//withotu pointers we would not be able to modify the value of the structure
	//we would only be able to modify a copy of the structure
	p.Age++
}

func main() {
	//let's create a person
	//we can create a person by creating a variable of type Person
	//and then assigning values to the fields of the structure
	//we can assign values to the fields of the structure using the dot operator
	//the dot operator is used to access fields of a structure

	//let's create a person called Alice
	alice := Person{
		Name: "Alice",
		Age:  25,
	}

	//we can create a person called Bob
	bob := Person{
		Name: "Bob",
	} //we will add age later

	//let's print these structures
	fmt.Println(alice)
	fmt.Println(bob)
	//let's modify Bob's age
	bob.Age = 30
	fmt.Println(bob)

	//let's make alice do the greeting
	alice.greet() //note how the syntax is different from calling a function
	//we use the dot operator to call a structure method
	//let's have alice have a birthday
	alice.birthday() //so in place mutation of the structure
	alice.greet()    //so alice is now 26

}