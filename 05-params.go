package tour

import "fmt"

// shared types
func add2(x, y int) int {
	return x + y
}

func call_add() {
	fmt.Println(add2(42, 13))
}

