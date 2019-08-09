package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/uesteibar/phenomena_calendar_scraper/scrape/calendar"
	"github.com/uesteibar/phenomena_calendar_scraper/scrape/phenomena"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	month := phenomena.FetchMonth(2019, 8)
	ics := calendar.CreateICS(month)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            ics,
		Headers: map[string]string{
			"Content-Type":        "text/calendar",
			"charset":             "utf-8",
			"Content-Disposition": "inline",
			"filename":            "phenomena_calendar.ics",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
