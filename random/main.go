package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	lambda.Start(Handler)
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	values, _ := url.ParseQuery(request.Body)

	parameters := make(map[string]string)
	for key := range values {
		parameters[key] = values[key][0]
	}

	if r, err := generateRandomInteger(parameters); err != nil {
		return events.APIGatewayProxyResponse{
			Body:       err.Error(),
			StatusCode: 400,
		}, nil
	} else {
		return events.APIGatewayProxyResponse{
			Body:       fmt.Sprintf(`{"text": "Your random number is: **%d**","response_type": "in_channel"}`, r),
			StatusCode: 200,
		}, nil
	}
}

func generateRandomInteger(queryStringParameters map[string]string) (int, error) {
	parameters := strings.Split(strings.Trim(queryStringParameters["text"], " "), " ")

	if len(parameters) == 1 {
		if parameters[0] == "" {
			return rand.Int() % 101, nil
		} else {
			if parameters[0] == "die" || parameters[0] == "dice" {
				return rand.Int()%6 + 1, nil

			} else if i, err := strconv.Atoi(parameters[0]); err != nil {
				return 0, fmt.Errorf("Unexpected parameter: %s", parameters[0])
			} else {
				return rand.Int() % (i + 1), nil
			}
		}
	} else if len(parameters) == 2 {
		i1, err := strconv.Atoi(parameters[0])
		if err != nil {
			return 0, fmt.Errorf("Unexpected parameter: %s", parameters[0])
		}

		i2, err := strconv.Atoi(parameters[1])
		if err != nil {
			return 0, fmt.Errorf("Unexpected parameter: %s", parameters[1])
		}

		return i1 + rand.Int()%(i2-i1+1), nil
	} else {
		return 0, fmt.Errorf("Unexpected # of parameters: %d", len(parameters))
	}
}
