package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func main() {
	result := plus(1, 3)
	fmt.Println(result)
	result = Plus2(1, 3)
	fmt.Println(result)
	var s int
	var t float32
	s = 123
	t = 2.3
	fmt.Println(s + int(t))

}
