package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	// Create & parse input flags
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffles the questions in the quiz")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines, *shuffle)

	// Setup timer
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

	for i, problem := range problems {
		fmt.Printf("Problem #%d: %s = \n", i+1, problem.q)

		// Setup a local go routine and set to answerCh
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		// What we do if time expires
		case <-timer.C:
			exitMessage(correct, len(problems))
			return
		// What we do if user submitted input
		case answer := <-answerCh:
			if answer == problem.a {
				correct++
			}
		}
	}

	exitMessage(correct, len(problems))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func exitMessage(correct int, total int) {
	fmt.Printf("You scored %d / %d\n", correct, total)
}

func parseLines(lines [][]string, shuffle bool) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}
	if shuffle {
		shuffleProblems(&ret)
	}

	return ret
}

// Fisher-Yates shuffle
func shuffleProblems(ordered *[]problem) {
	//shuffled := make([]problem, len(ordered))
	for i := len(*ordered) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		(*ordered)[i], (*ordered)[j] = (*ordered)[j], (*ordered)[i]
	}
}
