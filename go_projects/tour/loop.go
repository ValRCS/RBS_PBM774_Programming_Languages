package main

import "fmt"

//so project should generally have a single main package
//and the main package should have a single main function
//this is the entry point of the program
//for teaching we are compiling and running individual files so not a big deal
//but in a real project you would have to compile all the files in the package
//and the main function would be the entry point
//so in this case we would have to compile all the files in the tour package
func main() {
	sum := 0
	//so for loop is the only loop in Go
	//we have pre,  condition, and post
	//we can omit the pre and post
	//so this is equivalent to while loop in other languages
	//note that we don't have while loop in Go
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//lets imitate a while loop
	total := 0
	for total < 10 {
		//I could do whatever I want here inside while loop, technically a for loop
		//but I am just going to increment total
		total++
	}
	fmt.Println(total)
}