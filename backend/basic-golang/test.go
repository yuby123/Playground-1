package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		if i == 5 {
			break
		}
		if i == 2 {
			continue
		}
		fmt.Println(i)
	}
}
