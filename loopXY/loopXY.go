package loopxy

import "fmt"

func main() {
	var x, y, z int
	fmt.Scanln(&x)
	fmt.Scanln(&y)
	z = x
	y = y - 1
	for a := 0; a < y; a++ {
		z = z * x
	}
	fmt.Println(z)
}
