package main

import (
	"fmt"
	"math"
	"math/rand"
	"math/cmplx"
)

//A var statement can be at package or function level. We see both in this example.
// var c, python, java bool
var c, python, java = true, false, "no!" //so c should be true, python should be false and java should be "no!"
//note we use single assignment for all three variables
//also note that we can specify the type of the variables, but we don't have to, because the compiler can infer the type from the value

func add(x int, y int) int { //note how type is specified AFTER the variable name
	return x + y
}

//multiple results
func swap(x, y string) (string, string) {
	return y, x
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	//so there is some seed for random number generation
	// rand.Seed(2025) //deprecated
	//rand.Seed(time.Now().UnixNano()) //this will generate a different random number every time
	fmt.Println("My favorite number is", rand.Intn(10))
	//how about printing some random float?
	fmt.Println("Random Float is", rand.Float64())

	//now let's pring math.Pi
	fmt.Println("Math Pi is", math.Pi) //again note th capitalization

	//lets use the add function
	fmt.Println("The sum of 2 and 3 is", add(2, 3))

	a, b := swap("hello", "world") //so a should be "world" and b should be "hello"
	//also note the use of := instead of var this means that the type is inferred, so no need to specify it
	//so in this case means that a and b will have STATIC type string and their values will be "world" and "hello" respectively
	fmt.Println(a, b)

	var i int
	k := 3 //so k should be 3 
	//we can only use := inside a function, outside we have to use var
	//also note that we can't use := for reassignment, only for initialization
	fmt.Println(k, i, c, python, java) //so i should be 0 and the rest should be false, BECAUSE bool is false by default and uninitialized int is 0

	var (
		ToBe   bool       = false
		MaxInt uint64     = 1<<64 - 1 //this is a bit shift operation, so 1<<64 is 2^64 then -1 is 2^64 - 1
		z      complex128 = cmplx.Sqrt(-5 + 12i)
	)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)


	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var zz uint = uint(f)
	// fmt.Println(x, y, zz)
	//let's print Type and value of x, y, f and zz
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", f, f)
	fmt.Printf("Type: %T Value: %v\n", zz, zz)

	//we can also declare constants
	const Truth = true
	fmt.Println("Go rules?", Truth)

	//how about myPi?
	const myPi = 3.14 //the type of myPi is inferred to be float64
	//print type and value of myPi
	fmt.Printf("Type: %T Value: %v\n", myPi, myPi)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
