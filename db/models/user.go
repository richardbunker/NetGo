package models

import (
	"NetGo/db"
	NetGoTypes "NetGo/types"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"os"
)

func ListUsers() ([]NetGoTypes.User, bool) {
	inputModifiers := func(input *dynamodb.QueryInput) {
		input.KeyConditionExpression = aws.String("PK = :v1")
		input.ExpressionAttributeValues = map[string]types.AttributeValue{
			":v1": &types.AttributeValueMemberS{Value: "USER"},
		}
	}
	items, err := db.SingletonInstance.Query(inputModifiers)
	if err != nil {
		fmt.Println("Error querying DynamoDB: ", err)
	}
	if len(items) == 0 {
		return []NetGoTypes.User{}, false
	}
	var users []NetGoTypes.User
	for _, item := range items {
		users = append(users, NetGoTypes.User{
			Id:    item["SK"].(*types.AttributeValueMemberS).Value,
			Email: item["GSISK"].(*types.AttributeValueMemberS).Value,
			Name:  item["Name"].(*types.AttributeValueMemberS).Value,
		})
	}
	return users, true
}

// Find a user by email
func FindUserByEmail(email string) (NetGoTypes.User, bool) {
	inputModifiers := func(input *dynamodb.QueryInput) {
		input.KeyConditionExpression = aws.String("GSIPK = :v1 AND GSISK = :v2")
		input.ExpressionAttributeValues = map[string]types.AttributeValue{
			":v1": &types.AttributeValueMemberS{Value: "USER"},
			":v2": &types.AttributeValueMemberS{Value: email},
		}
		input.IndexName = aws.String(os.Getenv("DDB_GSI_NAME"))
	}
	items, err := db.SingletonInstance.Query(inputModifiers)
	if err != nil {
		fmt.Println("Error querying DynamoDB: ", err)
	}
	if len(items) == 0 {
		return NetGoTypes.User{}, false
	}
	item := items[0]
	return NetGoTypes.User{
		Id:    item["SK"].(*types.AttributeValueMemberS).Value,
		Email: item["GSISK"].(*types.AttributeValueMemberS).Value,
		Name:  item["Name"].(*types.AttributeValueMemberS).Value,
	}, true
}

// Find a user by id
func FindUserById(userId string) (NetGoTypes.User, bool) {
	inputModifiers := func(input *dynamodb.QueryInput) {
		input.KeyConditionExpression = aws.String("PK = :v1 AND SK = :v2")
		input.ExpressionAttributeValues = map[string]types.AttributeValue{
			":v1": &types.AttributeValueMemberS{Value: "USER"},
			":v2": &types.AttributeValueMemberS{Value: userId},
		}
	}
	items, err := db.SingletonInstance.Query(inputModifiers)
	if err != nil {
		fmt.Println("Error querying DynamoDB: ", err)
	}
	if len(items) == 0 {
		return NetGoTypes.User{}, false
	}
	item := items[0]
	return NetGoTypes.User{
		Id:    item["SK"].(*types.AttributeValueMemberS).Value,
		Email: item["GSISK"].(*types.AttributeValueMemberS).Value,
		Name:  item["Name"].(*types.AttributeValueMemberS).Value,
	}, true
}

// Register a new user
func RegisterUser(email string, name string) (NetGoTypes.User, error) {
	userId := uuid.New().String()
	user := NetGoTypes.UserDynamoDBItem{
		PK:    "USER",
		SK:    userId,
		GSIPK: "USER",
		GSISK: email,
		Name:  name,
		Type:  "User",
	}
	userItem, err := attributevalue.MarshalMap(user)
	if err != nil {
		panic(err)
	}
	inputModifiers := func(input *dynamodb.PutItemInput) {
		input.Item = userItem
	}
	err = db.SingletonInstance.AddItem(inputModifiers)
	if err != nil {
		return NetGoTypes.User{}, fmt.Errorf("failed to add item, %v", err)
	}
	return NetGoTypes.User{
		Id:    userId,
		Email: email,
		Name:  name,
	}, nil
}
