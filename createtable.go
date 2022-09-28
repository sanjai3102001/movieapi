package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// This is the function to CreateTable
func CreateTablee() {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})
	tableName := "Movies"
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Movieid"),
				AttributeType: aws.String("N"),
			},
			{
				AttributeName: aws.String("Title"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Movieid"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Title"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(tableName),
	}
	_, err := svc.CreateTable(input)
	if err != nil {
		log.Fatal("Got error calling CreateTable:", err)
	}
	fmt.Println("Created the table", tableName)
}
