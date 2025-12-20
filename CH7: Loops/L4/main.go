package main

import (
	"fmt"
)

func fizzbuzz() {
	for i := 0; i <= 100; i++ {
		var fizzbuzz string

		if i%3 == 0 {
			fizzbuzz += fizz()
		}

		if i%5 == 0 {
			fizzbuzz += buzz()
		}

		if fizzbuzz != "" {
			fmt.Println(fizzbuzz)
		} else {
			fmt.Println(i)
		}

	}
}

func fizz() string {
	return "fizz"
}

func buzz() string {
	return "buzz"
}

// don't touch below this line

func main() {
	fizzbuzz()
}
