package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	rightAnswers = 0
	rowsAmount   = 0
)

func runTimer(d time.Duration) {
	time.AfterFunc(d, func() {
		getResult("Time's up!", rightAnswers, rowsAmount)
		os.Exit(0)
	})
}

func main() {
	file, err := os.Open("test.csv")

	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading the lines", err)
	}

	rowsAmount = len(records)

	timerPointer := flag.Int("t", 10, "timer duration in seconds")
	flag.Parse()

	fmt.Println("press any key")
	readyInput := bufio.NewScanner(os.Stdin)
	readyInput.Scan()
	go runTimer(time.Duration(*timerPointer) * time.Second)
	rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })

	for _, record := range records {
		fmt.Println("Solve the quiz: " + record[0])
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if strings.TrimSpace(input.Text()) == record[1] {
			rightAnswers++
		}
	}
	getResult("You answered all questions", rightAnswers, rowsAmount)
}

func getResult(text string, rightAnswers int, rowsAmount int) {
	fmt.Printf("%s. You got %d out of %d right!", text, rightAnswers, rowsAmount)
}
