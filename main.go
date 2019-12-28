package main

import(
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Qst struct {
	question string
	ans string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Reads a file and returns a list of strings (lines of text)
func readFile(filename string) []Qst{
	csvFile, err := os.Open(filename)
	check(err)
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var questions []Qst
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		questions = append(questions, Qst{
			question: line[0],
			ans: line[1],
		})
	}
	return questions
}

func main() {
	fmt.Println("Welcome the Quiz!")
	questions := readFile("problems.csv")
	var correct, incorrect int
	for _, question := range questions {
		fmt.Println(question.question)
		userInput := ""
		_, err := fmt.Scanln(&userInput)
		check(err)
		if userInput == question.ans {
			correct++
		} else {
			incorrect++
		}
	}
	fmt.Printf("You got %v correct and %v incorrect.", correct, incorrect)
}