package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	q string
	a string
}

func main() {

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
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

	correct := 0
	for i, line := range data {
		fmt.Printf("Problem #%d: %s = ", i+1, line.q)
		var answer string
		fmt.Scanf("%s/n", &answer)
		if answer == line.a {
			correct++
			fmt.Println("Correct Answer ✅")
		} else {
			fmt.Println("Wrong Answer ❌")
		}
	}
	fmt.Printf("You Scored %d Out Of %d", correct, len(data))
}

func parseLines(lines [][]string) []problem {
	arr := make([]problem, len(lines))
	for index, line := range lines {
		arr[index] = problem{
			q: line[0],
			a: line[1],
		}
	}
	return arr
}

func exit(err string) {
	fmt.Println(err)
	os.Exit(1)
}
