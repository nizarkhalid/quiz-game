package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	csvFileName := flag.String("csv", "problems.csv", "a csv file that is formated as 'question,answer'")
	timerLimit := flag.Int("limit", 30, "a number that represent time in seconds")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("failed to open file %s", *csvFileName))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("couldn't read the file")
	}
	problems := parselines(lines)
	timer := time.NewTimer(time.Duration(*timerLimit) * time.Second)
	rightanswers, wronganswers := 0, 0
breakloop:
	for i, p := range problems {
		fmt.Printf("problem #%d: %s=", i+1, p.q)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s /n", &answer)
			answerCh <- answer
		}()
		select {
		case <-timer.C:
			fmt.Println()
			break breakloop
		case answer := <-answerCh:
			if answer == p.a {
				fmt.Println("correct!")
				rightanswers++
			} else {
				fmt.Println("wrong!!")
				wronganswers++
			}
		}

	}
	fmt.Printf("you made %d mistakes and scored %d points\n", wronganswers, rightanswers)
}

func parselines(lines [][]string) []problems {
	ret := make([]problems, len(lines))
	for i, line := range lines {
		ret[i] = problems{
			q: line[0],
			a: line[1],
		}
	}
	return ret
}

type problems struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
