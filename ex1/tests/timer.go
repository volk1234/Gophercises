package main

import "fmt"
import "time"

func main(){
    var input string
    timer := time.AfterFunc(time.Second*4, func() {
        fmt.Printf("you're out!")
		return 
    })
	fmt.Println(timer)
    defer timer.Stop()

    fmt.Println(">> ")
    _,err := fmt.Scanf("%s\n", &input)
	if err != nil {
		fmt.Printf("Error reading user input: %v\n",err)
	}else{
		timer.Stop()
	}

}