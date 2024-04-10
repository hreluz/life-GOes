package newmake

import "fmt"

func main() {
	// new & make

	fmt.Println("====== Using simplest way ======")
	hobbies := []string{"Sports", "Reading"}

	// go creates a new array behind the scenes, bc the old array is not big enough
	hobbies = append(hobbies, "Cooking", "Dancing")

	fmt.Println(hobbies)

	fmt.Println("====== Using make ======")
	// 1st arg  => type || 2nd  arg => length of the slice || 3rd arg => how big should be
	hobbies_make := make([]string, 2, 10)
	// aMap := make(map[string]int, 5)

	hobbies_make[0] = "Sports"
	hobbies_make[1] = "Reading"

	hobbies_make = append(hobbies_make, "Cooking", "Dancing")
	fmt.Println(hobbies_make)

	fmt.Println("====== Using new ======")
	// new returns a pointer
	hobbies_new := new([]string)

	fmt.Println(hobbies_new)
	fmt.Println(*hobbies_new)

	*hobbies_new = append(*hobbies_new, "Sports")

	fmt.Println(hobbies_new)
	fmt.Println(*hobbies_new)

	number := new(int)
	// address
	fmt.Println(number)
	// the number
	fmt.Println(*number)

	anotherNumber := 0
	numberAddress := &anotherNumber
	fmt.Println(numberAddress)
	fmt.Println(*numberAddress)
}
