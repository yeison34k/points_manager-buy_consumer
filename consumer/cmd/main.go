package main

import (
	"consumer/internal/adapter"
	"consumer/internal/app"
	"consumer/internal/infrastructure/sqs"
	usecase "consumer/internal/usercase"
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaHandler struct {
	pointHandler *app.BuyHandler
}

func NewLambdaHandler() *LambdaHandler {
	queueURL := os.Getenv("QUEUE_URL")
	if queueURL == "" {
		log.Fatal("La variable de entorno QUEUE_URL es requerida")
	}

	serviceURL := os.Getenv("SERVICE_URL")
	if queueURL == "" {
		log.Fatal("La variable de entorno SERVICE_URL es requerida")
	}

	httpClient := &adapter.HTTPAdapter{}

	httpClientCase := &usecase.HttpClientCase{
		HTTPClient: httpClient,
	}

	sqsHandler := sqs.NewSQSHandler(queueURL, serviceURL, httpClientCase)
	pointApp := app.NewBuyApplication(sqsHandler)
	pointHandler := app.NewBuyHandler(pointApp)
	return &LambdaHandler{
		pointHandler,
	}
}

func (h *LambdaHandler) HandleRequest(ctx context.Context, sqsEvent events.SQSEvent) error {
	err := h.pointHandler.HandleBuyCreation()
	if err != nil {
		log.Fatal("Error HandleBuyCreation:", err)
		return err
	}

	return nil
}

func main() {
	handler := NewLambdaHandler()
	lambda.Start(handler.HandleRequest)
}
