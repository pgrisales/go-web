package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"whats your name?"`
	Age  int    `json:"how old are you?"`
}

type Myresponse struct {
	Message string `json:"Answer"`
}

func HandleLambdaEvent(event MyEvent) (Myresponse, error) {
	return Myresponse{Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age)}, nil
}
func main() {
	lambda.Start(HandleLambdaEvent)
}
