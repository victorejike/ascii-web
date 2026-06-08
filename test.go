package main

import (
	"fmt"
)

func main() {
	result := [2]string{"victor", "ejike"}

	fmt.Println(result, result[0], result[len(result)-1])

	result[1] = "nmesomma"

	fmt.Println(result)
}
