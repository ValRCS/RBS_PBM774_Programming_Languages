//let's explore Arrays and Slices in Go language

package main

import "fmt"

func main() {
	//let's declare an array of 5 integers
	var arr [5]int
	fmt.Println("Empty Array:", arr)

	//let's declare an array with values
	arr1 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array with values:", arr1)

	//we can acess and mutate the values of an array using index
	fmt.Println("Value at index 2:", arr1[2])
	arr1[2] = 10
	fmt.Println("Array after mutation:", arr1)

	//now let's talk about a slice

	//slice in Go is flexible and dynamic view of an array
	//we can create a slice from an array or another slice
	//we can also create a slice using make function

	//let's create a slice from an array
	slice1 := arr1[1:4] // so again we take index 1 to 3 but not 4 (just like python)
	fmt.Println("Slice from an array:", slice1)
	//now when we modify say 2nd value in slice, it will also modify the original array
	slice1[1] = 20
	//print slice
	fmt.Println("Slice after modification:", slice1)
	//print original array
	fmt.Println("Array after modifying slice:", arr1)

	//slices have capacity and length
	//length is the number of elements in the slice
	//capacity is the number of elements in the underlying array starting from the first element in the slice
	//we can get the length and capacity of a slice using len and cap functions
	//let's see length and capacity of slice1
	fmt.Println("Length of slice1:", len(slice1))
	fmt.Println("Capacity of slice1:", cap(slice1))

	//how about length of an array
	fmt.Println("Length of arr1:", len(arr1))

	//how about we create a slice from beginning of an array for say 2 elements
	slice2 := arr1[:2] //so we take elements from beginning to index 1, very similar to python, 0 is implicit
	fmt.Println("Slice from beginning of an array:", slice2)
	//let's see length and capacity of slice2
	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2)) //this capacity is 5 as it is the capacity of underlying array,because we started from beginning of array

	//how about changing slice1 to have capacity of 5 as well?
	slice1 = append(slice1, 30, 40)
	fmt.Println("Slice1 after appending 30 and 40:", slice1)
	//let's see length and capacity of slice1
	fmt.Println("Length of slice1:", len(slice1))
	fmt.Println("Capacity of slice1:", cap(slice1)) //this capacity is 5 as it is the capacity of underlying array
	//let's print our underlying array
	fmt.Println("Array after appending to slice1:", arr1)
	//let's access the underlying array using slice1
	slice1[0] = 100
	fmt.Println("Array after modifying slice1:", arr1)

	//let's append two elements to slice2
	slice2 = append(slice2, 50, 60)
	fmt.Println("Slice2 after appending 50 and 60:", slice2)
	//let's see length and capacity of slice2
	fmt.Println("Length of slice2:", len(slice2))
	fmt.Println("Capacity of slice2:", cap(slice2))

	//let's see the underlying array
	fmt.Println("Array after appending to slice2:", arr1)

	//let's create an array of million elements
	var arr2 [1e6]int //so 1e6 is 10^6
	//let's populate the array with values from 0 to 999999
	for i := 0; i < 1e6; i++ { //of course I could have used 1000000 instead of 1e6
		arr2[i] = i
	}
	//let's create a slice from this array
	slice3 := arr2[:999999] //so we take elements from beginning to index 999998
	//let's see length and capacity of slice3
	fmt.Println("Length of slice3:", len(slice3))   //length should be 999999
	fmt.Println("Capacity of slice3:", cap(slice3)) //this capacity is 1000000 as it is the capacity of underlying array
	//length of original array
	fmt.Println("Length of arr2:", len(arr2))
	//lets print last 5 elements of slice3
	fmt.Println("Last 5 elements of slice3:", slice3[999994:])

	//let's append 3 values to our slice3
	slice3 = append(slice3, 999999, 1000000, 1000001)
	//let's see length and capacity of slice3
	fmt.Println("Length of slice3:", len(slice3))   //length should be 1000002
	fmt.Println("Capacity of slice3:", cap(slice3)) //this capacity is actually 1250304 not 2000000 in Go version go version go1.23.4 linux/amd64
	//length of original array
	fmt.Println("Length of arr2:", len(arr2))
	//lets print last 5 elements of slice3
	fmt.Println("Last 5 elements of slice3:", slice3[len(slice3)-5:]) //there is no negative indexing in Go, so we use len(slice3)-5 to get last 5 elements

	//let's copy a slice (meaning the underlying array is copied as well)
	slice4 := make([]int, len(slice3))
	copy(slice4, slice3)
	//let's see length and capacity of slice4
	fmt.Println("Length of slice4:", len(slice4))   //length should be 1000002
	fmt.Println("Capacity of slice4:", cap(slice4)) //this capacity should be something like 1.2 million something
	//let's see the last 5 elements of slice4
	fmt.Println("Last 5 elements of slice4:", slice4[len(slice4)-5:])
	//let's modify slice4
	slice4[0] = 1000002
	//let's see the first 5 elements of slice3
	fmt.Println("First 5 elements of slice3:", slice3[:5]) //should be 0, 1, 2, 3, 4
	//let's see the first 5 elements of slice4
	fmt.Println("First 5 elements of slice4:", slice4[:5]) //should be 1000002, 1, 2, 3, 4

	//TODO how to get the reference to array that lies underneath a slice?

}
