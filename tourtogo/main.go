package main

import (
	"tourtogo/Helper"
	"fmt"
)

func main() {
	UserName := Helper.GetString("what is your name ?")

	fmt.Printf("Hello, %s, it's great meeting you\n", UserName)

	favouriteColor := Helper.GetString("what is your favourite color?")
	fmt.Printf("wow, %s is a great color!\n", favouriteColor)
}
