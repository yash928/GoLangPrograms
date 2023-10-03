package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("hello")
	numb := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range numb {
		if num%2 == 0 {
			fmt.Println("The number is even", num)
		} else {
			fmt.Println("The number is odd", num)
		}
	}
	fmt.Println("Enter a number")
	reader := bufio.NewReader(os.Stdin)
	inp, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(inp)
	}
}
