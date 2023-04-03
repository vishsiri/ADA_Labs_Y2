//64050229
//Vish Siriwatana

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var numbers []int
	var nindex []int

	                             := 100

	for i := 0; i < 100; i++ {
		numbers = append(numbers, rand.Intn(numbersrange+1))
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
		if (i+1)%20 == 0 {
			fmt.Println()
		}
	}

	fmt.Print("Enter a number to search the index : ")
	reader := bufio.NewReader(os.Stdin)
	var input int
	fmt.Fscanln(reader, &input)

	//Linear Search Algorithm
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == input {
			nindex = append(nindex, i)
		}
	}

	//Print the index of the number in the array if array not empty else print not found
	if len(nindex) != 0 {
		fmt.Printf("Found number %d at index : ", input)
		for i := 0; i < len(nindex); i++ {
			fmt.Print(nindex[i], " ")
		}
	} else {
		fmt.Print("Not Found")
	}

}
