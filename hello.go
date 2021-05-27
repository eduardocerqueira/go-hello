package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"

	"example.com/user/hello/morestrings"
	"example.com/user/hello/utils"
	"github.com/google/go-cmp/cmp"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// naked return
	return
}

var name string = "Eduardo"
var i, j, k int = 1, 2, 3
var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	f1     float64
	f2     bool
	f3     string = "testing"
)

// constant
const Pi = 3.14

func main() {
	fmt.Println("the time is:", time.Now())
	fmt.Println("Hello, GO!")

	// packages
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	fmt.Println(rand.Intn(10))
	fmt.Println(math.Sqrt(7))
	fmt.Println(math.Pi)

	// functions
	fmt.Println(add(10, 20))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	x := 32
	fmt.Println("splitting", x)
	fmt.Println(split(x))

	// variables
	fmt.Println(name)
	fmt.Println(i, j, k)

	// vars inside functions with short statement declaration
	p := 3
	v, t := true, false
	fmt.Println(p, v, t)

	// vars type
	fmt.Printf("type: %T value: %v\n", ToBe, ToBe)
	fmt.Printf("type: %T value: %v\n", MaxInt, MaxInt)
	fmt.Printf("%v %v %v\n", f1, f2, f3)

	// constant
	fmt.Println(Pi)

	err := utils.WriteToFile("test.txt", "testing writing to a file!!!\n")
	if err != nil {
		log.Fatal(err)
	}

}
