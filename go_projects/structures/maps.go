//maps in Go programming language give us key value pairs
//maps are unordered
//maps are reference types
//maps are dynamic - we can add and remove elements from a map
//maps are used to represent a collection of elements where each element is stored as a key value pair
//access is constant time - O(1)
//delete is constant time - O(1)
//insert is constant time - O(1)
//search by key is constant time - O(1)
//maps are not thread safe - there is a sync package in Go that can be used to make maps thread safe

package main //goes without saying that your project should only have one package main and one main function

import (
	"fmt"
	"sort"
)

func main() {
	//let's create a map of names from "Alice", "Bob", "Carol", "Dave" and their ages
	ages := map[string]int{ //so keys are strings here and values are integers - remember static typing in Go
		"Alice": 25,
		"Bob":   30,
		"Carol": 35,
		"Dave":  40,
	}

	//print the map
	fmt.Println(ages)

	//we can get values from a map using the key
	fmt.Println(ages["Alice"]) //25
	//we can delete a key value pair from a map
	delete(ages, "Bob")
	fmt.Println(ages) //Bob is gone

	//we can also add a new key value pair to a map let's add "Eve" to the map
	ages["Eve"] = 45
	fmt.Println(ages) //Eve is added

	//we can also check if a key exists in a map
	//if a key exists in a map, the value returned is the value of the key
	//if a key does not exist in a map, the value returned is the zero value of the value type
	//so if we check for a key that does not exist in a map, we get the zero value of the value type

	//let's check if "Alice" exists in the map
	age, ok := ages["Alice"] //ok is a boolean that tells us if the key exists in the map
	if ok {
		fmt.Println("Alice exists in the map and her age is", age)
	} else {
		fmt.Println("Alice does not exist in the map")
	}

	//soo above is a very typical on how possible error handling is done in Go

	//let's iterate over the map
	for key, value := range ages {
		fmt.Println(key, value)
	}
	//note how order is not guaranteed in the map!!
	//each build and run gives us a different order!

	//so if we want to print key value pairs in order
	//we would have to sort the keys first
	//let's sort the keys first
	//we can use the sort package in Go to sort the keys
	//first let's get keys
	var keys []string
	for key := range ages {
		keys = append(keys, key)
	}
	//now let's sort the keys
	sort.Strings(keys) //so keys are sorted IN PLACE
	//now let's print the keys in order
	for _, key := range keys { //note how we skipped the value here
		fmt.Println(key, ages[key])
	}
}
