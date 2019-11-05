package repository

import (
	"github.com/sslab-instapay/instapay-tee-client/model"
	"github.com/sslab-instapay/instapay-tee-client/db"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func InsertPaymentData(paymentData model.PaymentData) (model.PaymentData, error) {
	database, err := db.GetDatabase()
	if err != nil {
		return model.PaymentData{}, err
	}

	collection := database.Collection("payments")

	insertResult, err := collection.InsertOne(context.TODO(), paymentData)
	if err != nil {
		return model.PaymentData{}, err
	}

	fmt.Println(insertResult.InsertedID)

	return paymentData, nil
}

func FindPaymentData(data model.PaymentData) (bool, error){
	database, err := db.GetDatabase()
	if err != nil {
		return false, err
	}

	filter := bson.M{
		"paymentNumber": data.PaymentNumber,
		"channelId": data.ChannelId,
		"amount": data.Amount,
	}
	collection := database.Collection("payments")
	singleInstance := collection.FindOne(context.TODO(), filter)
	var paymentData = model.PaymentData{}
	err = singleInstance.Decode(&paymentData)
	if err != nil{
		return false, err
	}
	return true, nil
}

func GetPaymentDatasByPaymentNumber(paymentNumber int64) ([]model.PaymentData, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"paymentNumber": paymentNumber,
	}

	collection := database.Collection("payments")
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	var paymentDatas []model.PaymentData

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var paymentData model.PaymentData
		err := cur.Decode(&paymentData)
		if err != nil {
			log.Println("Decode Error")
		}
		paymentDatas = append(paymentDatas, paymentData)
	}

	return paymentDatas, nil

}
