package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/mursalin-sk/gophercises/internal/quiz"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	csvFileName := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 8, "quiz time duration in second")
	flag.Parse()
	if csvFileName == nil {
		slog.Error("csv file name is missing")
		return
	}
	data, err := quiz.GetQuizDataFromCSV(*csvFileName)
	if err != nil {
		slog.Error("failed to parse quiz ", err)
		return
	}
	result := quiz.PlayQuiz(data, *timeLimit)
	fmt.Printf("You got %d out of %d", result, len(data))
}
