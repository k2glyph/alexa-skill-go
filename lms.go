package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

// Alexa Request
type AlexaRequest struct {
	Version string `json: "version"`
	Request struct {
		Type   string `json: "type"`
		Time   string `json: "timestamp"`
		Intent struct {
			Name               string `json: "name"`
			ConfirmationStatus string `json: "confirmationstatus"`
		} `json: "intent"`
	} `json: "request"`
}

// Alexa Response
type AlexaResponse struct {
	Version  string `json:"version"`
	Response struct {
		OutputSpeech struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"outputSpeech"`
	} `json:"response"`
}

// Create Response
func CreateResponse() *AlexaResponse {
	var resp AlexaResponse
	resp.Version = "1.0"
	resp.Response.OutputSpeech.Type = "PlainText"
	resp.Response.OutputSpeech.Text = "Hello.  Please override this default output."
	return &resp
}

// Alexa Say
func (resp *AlexaResponse) Say(text string) {
	resp.Response.OutputSpeech.Text = text
}

// Alexa HandleRequest
func HandleRequest(ctx context.Context, i AlexaRequest) (AlexaResponse, error) {
	fmt.Println(i)
	log.Printf("Request type is ", i.Request.Intent.Name)
	resp := CreateResponse()

	switch i.Request.Intent.Name {
	case "about":
		resp.Say("Auzmor LMS is simplified employee training software solution. Which is actively on development. For more information visit https://auzmor.com")
	case "hello":
		resp.Say("Hello there, Lambda appears to be working properly.")
	case "AMAZON.HelpIntent":
		resp.Say("This app is easy to use, just say: hello")
	default:
		resp.Say("I'm sorry, the input does not look like something I understand.")
	}
	return *resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
