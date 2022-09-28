package function

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Movieid int
	Title   string
	Hero    string
}

func ReadingItem(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	tableName := "Movies"
	Title := "kgf"
	movieid := "2010"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(movieid),
			},
			"Title": {
				S: aws.String(Title),
			},
		},
	})
	fmt.Println(result)
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	// id, err := strconv.Atoi(movieid)

	// if id != 0 {
	// 	log.Fatalf("No item: %s", err)
	// }

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Id:  ", item.Movieid)
	fmt.Println("Title: ", item.Title)
	fmt.Println("Hero:  ", item.Hero)

}

func ReadingItemid(w http.ResponseWriter, r *http.Request) {
	sess, _ := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("AKIASA45Q7S6M3LEBBL6", "SAJ95tqB1E6QTm0OMa5bUwS2vdm5tIYz2A/P9MZV", ""),
	})
	fmt.Println(sess.Config.Credentials.Get())
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("")})

	tableName := "Movies"
	Title := "kgf3"
	movieid := "2010"

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"Movieid": {
				N: aws.String(movieid),
			},
			"Title": {
				S: aws.String(Title),
			},
		},
	})
	fmt.Println(result)
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	//id, err := strconv.Atoi(movieid)

	if Title == "" {
		log.Fatalf("No item: %s", err)
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Println("Found item:")
	fmt.Println("Id:  ", item.Movieid)
	fmt.Println("Title: ", item.Title)
	fmt.Println("Hero:  ", item.Hero)

}
