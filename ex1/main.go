package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"flag"
)

func main (){

	qfile_name := flag.String("f", "/home/teth/repos/Gophercises/ex1/problems.csv", "path to quiz db file")
	flag.Parse()

	fmt.Println("file to load: ", *qfile_name)

	qfile, err := os.Open(*qfile_name)
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

	for i, record := range qrecords{
		if len(record) >=2 {
			var input int
			q := record[0]
			a := record[1]

			fmt.Printf("Question #%d: %s\n", i, q)
			answ, err := fmt.Scanf("%d\n", &input)
            if err != nil {
                fmt.Printf("error user input",err)
            }
			if answ == a {
				fmt.Println("Correct! ", answ, "equal", a)
			}else{
				fmt.Println("You are qrong! ", answ, "NOT equal", a)
			}
		}else{
            fmt.Println("Skipping malformed record:", record)
        }
	}
}