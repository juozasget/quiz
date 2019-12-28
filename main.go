package main

import(
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"
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
	userInput := ""
	var filename = flag.String("file", "problems.csv", "The problems csv file")
	var timeSeconds = flag.Int("time", 30, "Timer value in seconds")
	flag.Parse()

	fmt.Println("Welcome the Quiz!")
	questions := readFile(*filename)
	fmt.Println(*filename)
	fmt.Println("Press enter to start [ENTER]")
	fmt.Scanln(&userInput)
	timer := time.NewTimer(time.Duration(*timeSeconds) * time.Second)
	var correct int
	for i, question := range questions {
		fmt.Printf("Problem #%v: %v\n", i+1, question.question)
		answerCh := make(chan string)
		go func() {
			var userInput string
			_, err := fmt.Scanln(&userInput)
			check(err)
			answerCh <- userInput
		}()

		select {
		case <-timer.C:
			fmt.Printf("Time expired! You got %v correct out of %v.\n", correct, len(questions))
			return
		case userInput := <-answerCh:
			if userInput == question.ans {
				correct++
			}
		}
	}
	fmt.Printf("You got %v correct out of %v.\n", correct, len(questions))
}