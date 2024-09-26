package main

import (
	"fmt"
	"math/rand"
)

var arr = rand.Perm(200) //[6]int{3, 2, 8, 6, 1, 4}

func main()  {
	for i := 1; i < len(arr); i++ {
		fmt.Println("Next: ", arr[i])
		for j := 0; j < i; j++ {

			if arr[i] < arr[j]{
				fmt.Printf("%v\t",arr[j:i+1])
				fmt.Printf("Swap: %d with %d\t", arr[j], arr[i])
				arr[i], arr[j] = arr[j], arr[i]
				fmt.Println(arr[j:i+1])
				fmt.Printf("%v \n\n", arr)
			} else {
				fmt.Println("-= NO CHANGES =-")
			}
		}
	}
	fmt.Println(arr)
}
