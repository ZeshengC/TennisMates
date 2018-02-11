package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Request
type Request struct {
}


// Response
type Response struct {
	StatusCode float64 `json:"statusCode"`
	Headers    Headers 	`json:"headers"`
	Body	   string 	`json:"body"`
	IsBase64Encoded bool `json:"isBase64Encoded"`
}

// TennisCourtCoordicatesOnly
type TennisCourtCoordicatesOnly struct {
	ID int `json:"id"`
	Coordinates Coordinates `json:"coordinates"`
}


// Coordinates
type Coordinates struct{
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Headers
type Headers struct {
	ContentType string `json:"Content-Type"`
}

// toJSON
func toJson(p interface{}) string{
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}


// GetTennisCourts
func GetTennisCourts()[]TennisCourtCoordicatesOnly {
	raw, err := ioutil.ReadFile("courts.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	var t []TennisCourtCoordicatesOnly
	json.Unmarshal(raw,&t)
	return t
}

// Handler
func Handler(request Request)(Response, error){
	b := GetTennisCourts()
	r := Response{
		StatusCode: 200,
		Headers: Headers{ContentType:"application/json"},
		Body: toJson(b),
		IsBase64Encoded: false,
	}
	return r,nil
}

func main() {
	lambda.Start(Handler)
}