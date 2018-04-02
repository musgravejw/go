package tour

import "fmt"

func add1(x int, y int) int {
	return x + y
}

func sum() {
	fmt.Println(add1(42, 13))
}

