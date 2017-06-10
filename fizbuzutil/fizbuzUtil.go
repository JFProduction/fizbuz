package fizbuz

import (
	"fmt"
)

// FizBuz checks if the int passed
// should print fiz, buz or fizbuz,
// up to the value n
func FizBuz(n int) {
	if n <= 0 {
		fmt.Println("Please choose a number greater than 0")
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%v %v\n", i, fizbuz(i))
	}
}
