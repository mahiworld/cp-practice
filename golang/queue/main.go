package main

import "fmt"

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))    // Output: 1
	fmt.Println(obj.Ping(100))  // Output: 2
	fmt.Println(obj.Ping(3001)) // Output: 3
	fmt.Println(obj.Ping(3002)) // Output: 3
}
