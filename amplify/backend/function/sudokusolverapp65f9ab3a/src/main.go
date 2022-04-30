package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type SudokuEvent struct {
	Board [][]rune `json:"board"`
}

func HandleRequest(ctx context.Context, sudokuEvent SudokuEvent) (string, error) {
	return fmt.Sprintf("Hello %c!", sudokuEvent.Board[0][0]), nil
}

func main() {
	lambda.Start(HandleRequest)
}
