package main

//https://www.youtube.com/watch?v=NxntmGW62Ag&list=PLCKpcjBB_VlBsxJ9IseNxFllf-UFEXOdg&index=34

//Cap. 5 – Exercícios: Nível #2 – 4

import "fmt"

func main() {
	x := 200

	fmt.Printf("Binary %b\t\n", x)
	fmt.Printf("Hex %#x\t\n", x)
	fmt.Printf("Dec %d\t\n", x)
	fmt.Printf("_________________\n\n")
	y := x << 1

	fmt.Printf("Binary %b\t\n", y)
	fmt.Printf("Hex %#x\t\n", y)
	fmt.Printf("Dec %d\t\n", y)
}
