package db

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"os"
	"sync"
)

var (
	SingletonInstance *DynamoDBService
	once              sync.Once
)

type DynamoDBService struct {
	client *dynamodb.Client
}

func NewDynamoDBSingleton() error {
	var err error
	once.Do(func() {
		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(os.Getenv("AWS_REGION")))
		if err != nil {
			err = fmt.Errorf("failed to load configuration, %v", err)
			return
		}

		// Using the Config value, create the DynamoDB client
		svc := dynamodb.NewFromConfig(cfg)

		SingletonInstance = &DynamoDBService{
			client: svc,
		}

	})
	return err
}

func (d *DynamoDBService) AddItem(inputModifiers func(input *dynamodb.PutItemInput)) error {
	// Initialize PutItemInput with the table name
	input := &dynamodb.PutItemInput{
		TableName: aws.String(os.Getenv("DDB_TABLE_NAME")),
	}

	// Apply any additional modifications to the input
	inputModifiers(input)

	// Perform the PutItem operation
	_, err := d.client.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to put item, %v", err)
	}

	return nil
}

func (d *DynamoDBService) DeleteItem(inputModifiers func(input *dynamodb.DeleteItemInput)) error {
	// Initialize DeleteItemInput with the table name
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(os.Getenv("DDB_TABLE_NAME")),
	}

	// Apply any additional modifications to the input
	inputModifiers(input)

	// Perform the DeleteItem operation
	_, err := d.client.DeleteItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("failed to delete item, %v", err)
	}

	return nil
}

func (d *DynamoDBService) Query(inputModifiers func(input *dynamodb.QueryInput)) ([]map[string]types.AttributeValue, error) {
	// Initialize QueryInput with the table name
	input := &dynamodb.QueryInput{
		TableName: aws.String(os.Getenv("DDB_TABLE_NAME")),
	}

	// Apply any additional modifications to the input
	inputModifiers(input)

	// Perform the query operation
	resp, err := d.client.Query(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to query DynamoDB, %v", err)
	}

	// Check if items were returned
	if len(resp.Items) == 0 {
		return nil, nil
	}

	// Return and deference the items
	return resp.Items, nil
}
