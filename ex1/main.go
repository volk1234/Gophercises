package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"flag"
	"strings"
)

type quiz struct{
	q string
	a string
}

func coreQuiz (records []quiz){

    total := len(records)
	fmt.Printf("total is #%d:\n", total)
	correct := 0
	for i, record := range records{
			var input string
			fmt.Printf("Question #%d: %s\n", i+1, record.q)
			_, err := fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Printf("Error reading user input: %v\n",err)
			}

			if input == record.a {
				fmt.Printf("Correct! %s equals %s\n", input, record.a)
				correct++
			}else{
				fmt.Printf("Incorrect! %s does not equal %s\n", input, record.a)
			}

		fmt.Printf("Left: %d\n\n", total-i-1 )
	}
	fmt.Printf("Quiz results:\nTotal questions: %d\nCorrect answers:%d\nIncorrect answers:%d\n", total, correct, total-correct)
}

func main (){

	qFileName := flag.String("f", "problems.csv", "path to quiz db file")
	qTime := flag.Uint("t", 30, "time limit per question, sec.")
	flag.Parse()

	fmt.Println("file to load: ", *qFileName)
	fmt.Println("seconds to answer: ", *qTime)

	qfile, err := os.Open(*qFileName)
	if err != nil{
		dead(err.Error())
	}

	defer qfile.Close()

	c_reader := csv.NewReader(qfile)
	qrecords, err := c_reader.ReadAll()
	if err != nil{
		dead(err.Error())
	}

	questions := parseLines(qrecords)
	coreQuiz (questions)
}

func parseLines(lines [][]string) []quiz {
	ret := make([]quiz, len(lines))
	for i, line := range lines{
		ret[i] = quiz{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func dead(msg string){
	fmt.Printf(">> Fatal: %s\n", msg)
	os.Exit(1)
}