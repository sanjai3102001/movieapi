package function

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"log"
)

func UpdateItems(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	tableName := "Movies"
	Title := "kgf2"
	Movieid := "2010"
	Hero := "yash"

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":r": {
				S: aws.String(Hero),
			},
		},
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(Movieid),
			},
			"Title": {
				S: aws.String(Title),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set Hero = :r"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

	fmt.Println("Successfully updated '" + Title + "' (" + Movieid + ") rating to " + Title)

}
