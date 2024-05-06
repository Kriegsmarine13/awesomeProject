package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

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

	rowsAmount := len(records)
	rightAnswers := 0

	flag.Func("t", "Set timer for Quiz", func(s string) error {
		duration, err := time.ParseDuration(s)

		if err != nil {
			fmt.Println("Error in timer flag, using default timer 5")
			return err
		}

		time.NewTimer(duration)
		return nil
	})

	for _, record := range records {
		fmt.Println("Solve the quiz: " + record[0])
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == record[1] {
			rightAnswers++
		}
	}
	fmt.Println("Quiz is over. You got ", rightAnswers, " out of ", rowsAmount, " questions right!")

}
