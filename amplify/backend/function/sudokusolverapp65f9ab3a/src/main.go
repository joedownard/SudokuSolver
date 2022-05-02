package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type SudokuEvent struct {
	Board [9][9]int `json:"board"`
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var inputBoard SudokuEvent
	err := json.Unmarshal([]byte(request.Body), &inputBoard)
	if err != nil {
		log.Fatalf("unable to unmarshal request body, %v", err)
	}

	solvedBoardResponse := SudokuEvent{
		Board: SolveBoard(inputBoard.Board),
	}

	res, err := json.Marshal(solvedBoardResponse)
	if err != nil {
		log.Fatalf("unable to marshal board, %v", err)
	}

	return events.APIGatewayProxyResponse{
		Body:       string(res),
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
