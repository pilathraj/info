package main

import "fmt"

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcd(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	fmt.Println(gcd(24, 36), lcd(24, 36))

}
