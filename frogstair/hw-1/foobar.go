package frogstair

import "fmt"

// FooBar prints a foobar sequence until x
func FooBar(x int) {
	for i := 1; i <= x; i++ {
		if i%3 == 0 {
			fmt.Print("foo")
		}
		if i%5 == 0 {
			fmt.Print("bar")
		}
		if i%3 != 0 && i%5 != 0 {
			fmt.Print(i)
		}
		fmt.Println()
	}
}
