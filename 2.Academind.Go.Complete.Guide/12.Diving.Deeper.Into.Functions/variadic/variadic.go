package variadic

import (
	"fmt"
)

func main() {

	numbers := []int{1, 10, 25, 100}

	sum := sumup(1, 10, 25)

	fmt.Println(sum)

	anotherSum := sumup(1, numbers...)
	fmt.Println(anotherSum)

}

func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return startingValue + sum
}
