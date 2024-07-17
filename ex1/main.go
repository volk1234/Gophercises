package main

import (
	"fmt"
	"os"
	"encoding/csv"
	"flag"
	"strings"
	"time"
	"math/rand"
)

type quiz struct{
	q string
	a string
}

type param struct{
	filename *string
	timeLimit *uint
	shuffle *bool
}

func coreQuiz (records []quiz, duration uint){

    total := len(records)
	fmt.Printf("total is #%d:\n", total)
	timer := time.NewTimer(time.Duration(duration) * time.Second)

	correct := 0
	for i, record := range records{

		start := time.Now()
		fmt.Printf("Question #%d: %s\n", i+1, record.q)

		aCh := make(chan string)
		go func(){
			var input string
			_, err := fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Printf("Error reading user input: %v\n",err)
			}
			inputTrim := strings.TrimSpace(input)
			aCh <- inputTrim
		}()
		select {
		case <-timer.C:
			fmt.Println("You are out!")
			fmt.Printf("Quiz results:\nTotal questions: %d\nCorrect answers:%d\nIncorrect answers:%d\n", total, correct, total-correct)
			return
		case answer := <-aCh:

			if answer == record.a {
				fmt.Printf("Correct! %s equals %s\n", answer, record.a)
				t := time.Now()
				fmt.Println("Time elapsed", t.Sub(start))
				correct++
			}else{
				fmt.Printf("Incorrect! %s does not equal %s\n", answer, record.a)
			}

		fmt.Printf("Left: %d\n\n", total-i-1 )
		}
	}
		fmt.Printf("Quiz results:\nTotal questions: %d\nCorrect answers:%d\nIncorrect answers:%d\n", total, correct, total-correct)

}

func main (){

	param := param{}
	param.filename = flag.String("f", "problems.csv", "path to quiz db file")
	param.timeLimit = flag.Uint("t", 10, "time limit per question, sec.")
	param.shuffle = flag.Bool("s", true, "if shuffle quiz questions")
	flag.Parse()

	fmt.Println("file to load: ", *param.filename)
	fmt.Println("seconds to answer: ", *param.timeLimit)
	fmt.Println("shuffle questions: ", *param.shuffle)

	qfile, err := os.Open(*param.filename)
	if err != nil{
		dead(err.Error())
	}
	defer qfile.Close()
	
	c_reader := csv.NewReader(qfile)
	qrecords, err := c_reader.ReadAll()
	if err != nil{
		dead(err.Error())
	}
	questions := parseLines(qrecords, *param.shuffle)
	coreQuiz (questions,*param.timeLimit)
}

func parseLines(lines [][]string, shuffle bool) []quiz {
	ret := make([]quiz, len(lines))
	for i, line := range lines{
		ret[i] = quiz{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
		if shuffle == true {
		j := rand.Intn(i + 1)
		ret[i], ret[j] = ret[j], ret[i]
		}
	}
	return ret
}

func dead(msg string){
	fmt.Printf(">> Fatal: %s\n", msg)
	os.Exit(1)
}