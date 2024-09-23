package models

import (
	"NetGo/db"
	NetGoTypes "NetGo/types"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func FindLoginToken(token string) (NetGoTypes.LoginToken, error) {
	inputModifiers := func(input *dynamodb.QueryInput) {
		input.KeyConditionExpression = aws.String("PK = :v1")
		input.ExpressionAttributeValues = map[string]types.AttributeValue{
			":v1": &types.AttributeValueMemberS{Value: "LOGINTOKEN#" + token},
		}
	}
	items, err := db.SingletonInstance.Query(inputModifiers)
	if err != nil {
		fmt.Println("Error querying DynamoDB: ", err)
	}
	if len(items) == 0 {
		return NetGoTypes.LoginToken{}, fmt.Errorf("Login token not found")
	}
	item := items[0]
	return NetGoTypes.LoginToken{
		Id:        item["PK"].(*types.AttributeValueMemberS).Value,
		UserId:    item["SK"].(*types.AttributeValueMemberS).Value,
		ExpiresAt: item["ExpiresAt"].(*types.AttributeValueMemberS).Value,
	}, nil
}

func StoreLoginToken(userId string, token string, expiresAt string) error {
	loginTokenItem := NetGoTypes.LoginTokenDynamoDBItem{
		PK:        "LOGINTOKEN#" + token,
		SK:        userId,
		GSIPK:     "LOGINTOKEN",
		Type:      "LoginToken",
		ExpiresAt: expiresAt,
	}
	item, err := attributevalue.MarshalMap(loginTokenItem)
	if err != nil {
		panic(err)
	}
	inputModifiers := func(input *dynamodb.PutItemInput) {
		input.Item = item
	}
	err = db.SingletonInstance.AddItem(inputModifiers)
	if err != nil {
		return fmt.Errorf("failed to add item, %v", err)
	}
	return nil
}

func DeleteLoginToken(loginToken NetGoTypes.LoginToken) error {
	inputModifiers := func(input *dynamodb.DeleteItemInput) {
		input.Key = map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: loginToken.Id},
			"SK": &types.AttributeValueMemberS{Value: loginToken.UserId},
		}
	}
	err := db.SingletonInstance.DeleteItem(inputModifiers)
	if err != nil {
		return fmt.Errorf("failed to delete item, %v", err)
	}
	return nil
}
