package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timerValue := flag.Int("timer", 5, "timer for complete the quiz")
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the csv file : %s\n : with err %s", *csvFileName, err.Error()))
	}

	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		exit(" there an error when try to parse the file")
	}
	data := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timerValue) * time.Second)

	correct := 0
	for i, line := range data {
		fmt.Printf("Problem #%d: %s = ", i+1, line.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou Scored %d Out Of %d", correct, len(data))
			os.Exit(1)
		case answerValue := <-answerCh:
			if answerValue == line.a {
				correct++
			}
		}

	}
	fmt.Printf("You Scored %d Out Of %d", correct, len(data))
}

func parseLines(lines [][]string) []problem {
	arr := make([]problem, len(lines))
	for index, line := range lines {
		arr[index] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return arr
}

func exit(err string) {
	fmt.Println(err)
	os.Exit(1)
}
