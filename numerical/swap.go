package numerical

import "fmt"

func swap(x, y *int) {
	*x, *y = *y, *x
}

func gotest() {
	x := 3
	y := 2
	fmt.Println(x, y)
	swap(&x, &y)
	fmt.Println(x, y)
}
