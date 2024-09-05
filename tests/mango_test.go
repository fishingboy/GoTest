package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestFindOne(t *testing.T) {
	uri := "mongodb://127.0.0.1:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	assert.Nil(t, err)

	coll := client.Database("coffeeagent").Collection("user")
	userName := "master"

	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"name", userName}}).Decode(&result)
	assert.Nil(t, err)
	assert.Equal(t, "master", result["name"])

	jsonData, err := json.MarshalIndent(result, "", "    ")
	assert.Nil(t, err)
	fmt.Printf("%s\n", jsonData)
}

func TestUpdateOne(t *testing.T) {
	uri := "mongodb://127.0.0.1:27017/"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	assert.Nil(t, err)

	coll := client.Database("coffeeagent").Collection("user")
	userName := "master"

	result, err := coll.UpdateOne(
		context.TODO(),
		bson.D{{"name", userName}},
		bson.D{{"$set", bson.D{{"leo-age", 28}}}},
	)
	assert.Nil(t, err)
	fmt.Printf("The number of modified documents: %d\n", result.ModifiedCount)

	var userResult bson.M
	coll.FindOne(context.TODO(), bson.D{{"name", userName}}).Decode(&userResult)
	fmt.Println(userResult)
	var expected int32 = 28
	assert.Equal(t, expected, userResult["age"])
}
