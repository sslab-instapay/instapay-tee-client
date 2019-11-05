package repository

import (
	"github.com/sslab-instapay/instapay-tee-client/model"
	"github.com/sslab-instapay/instapay-tee-client/db"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"context"
)

func GetAllDepositValue() (int64, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return 0, err
	}

	filter := bson.M{"channelStatus": bson.M{
		"$not": bson.M{
			"$eq": model.CLOSED,
		},
	}, "channelType": bson.M{
		"$eq": model.IN,
	}}
	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return 0, err
	}
	var depositValue int64

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Println(err)
		}
		depositValue += channel.MyDeposit
	}

	return depositValue, nil
}

func GetOffChainBalance() (int64, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return 0, err
	}

	filter := bson.M{"channelStatus": bson.M{
		"$not": bson.M{
			"$eq": model.CLOSED,
		},
	}}
	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return 0, err
	}
	var myBalance int64

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Println(err)
		}
		myBalance += channel.MyBalance
	}

	return myBalance, nil
}
