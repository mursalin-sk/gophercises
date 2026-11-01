package quiz

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"os"
	"path"
	"time"

	"github.com/mursalin-sk/gophercises/pkg/models"
)

func GetQuizDataFromCSV(fileName string) ([]models.Quiz, error) {
	filePath := path.Join("data", "quiz", fileName)
	file, err := os.Open(filePath)
	if err != nil {
		slog.Error("failed to open csv file ", filePath)
		return nil, err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			slog.Error("failed to close csv file ", err)
		}
	}(file)
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		slog.Error("failed to open csv content", filePath)
		return nil, err
	}
	data := make([]models.Quiz, len(lines))
	for i, line := range lines {
		if len(line) == 2 {
			data[i] = models.Quiz{Question: line[0], Answer: line[1]}
		}
	}
	return data, nil
}

func PlayQuiz(data []models.Quiz, timeLimit int) int {
	correct := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	defer timer.Stop()

	for i, quiz := range data {
		fmt.Printf("Problem #%d: %s = ", i+1, quiz.Question)

		answerChan := make(chan string, 1)

		go func() {
			var answer string
			fmt.Scanf("%s", &answer)
			answerChan <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's up!")
			return correct

		case answer := <-answerChan:
			if answer == quiz.Answer {
				correct++
			}
		}
	}

	return correct
}
