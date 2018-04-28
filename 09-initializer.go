package tour

import "fmt"

var i, j int = 1, 2

func init() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
