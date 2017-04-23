package main

import "fmt"

func main() {
	fmt.Println(42, "\n")
	fmt.Printf("%d = %b \n", 42, 42)
	fmt.Printf("%d = %b = %x \n", 163, 163, 163)
	fmt.Printf("%d = %b = %x \n", 163, 163, 163)
	fmt.Printf("%d = %b = %#x \n", 163, 163, 163)
	fmt.Printf("%d = %b = %#X \n", 163, 163, 163)
	fmt.Printf("%d \t %b \t  %#X \n", 163, 163, 163)
	for i := 0; i < 100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
	for i := 0; i < 50; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}

}
