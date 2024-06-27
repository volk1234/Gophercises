package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"rsc.io/quote"
)

func main (){
	qfile, err := os.Open("/home/teth/repos/Gophercises/ex1/problems.csv")
	if err != nil{
		fmt.Println("File open error:", err)
		return
	}
	defer qfile.Close()

	creader := csv.NewReader(qfile)

	qrecords, err := creader.ReadAll()
	if err != nil{
		fmt.Println("CSV-file Read error: ", err)
		return
	}

	fmt.Println(quote.Glass())

	for _, record := range qrecords{
		if len(record) >=2 {
			q := record[0]
			a := record[1]

			fmt.Printf("Question: %s\nAnswer:%s\n", q, a)
		}else{
            fmt.Println("Skipping malformed record:", record)
        }
	}
}