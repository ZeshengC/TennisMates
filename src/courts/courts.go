package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Request struct {
}

type Response struct {
	StatusCode float64 `json:"statusCode"`
	Headers    Headers 	`json:"headers"`
	Body	   string 	`json:"body"`
	IsBase64Encoded bool `json:"isBase64Encoded"`
}

type TennisCourtCoordicatesOnly struct {
	ID int `json:"id"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct{
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type Headers struct {
	ContentType string `json:"Content-Type"`
}

func toJson(p interface{}) string{
	bytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return string(bytes)
}

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