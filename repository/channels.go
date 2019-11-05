package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"github.com/sslab-instapay/instapay-tee-client/model"
	"github.com/sslab-instapay/instapay-tee-client/db"
	"fmt"
	"sync"
)

func GetChannelList() ([]model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}
	var channels []model.Channel

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Println(err)
		}
		// To get the raw bson bytes use cursor.Current
		channels = append(channels, channel)
	}

	return channels, nil
}

func GetAllChannelsLockedBalance() (int64, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return 0, err
	}

	filter := bson.M{"channelStatus": bson.M{
		"$not": bson.M{
			"$eq": 3,
		},
	}}
	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return 0, err
	}
	var lockedBalance int64

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Println(err)
		}
		lockedBalance += channel.LockedBalance
	}

	return lockedBalance, nil
}

func GetChannelIdList() ([]int64, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		return nil, err
	}
	var channelIds []int64

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Println(err)
		}
		// To get the raw bson bytes use cursor.Current
		channelIds = append(channelIds, channel.ChannelId)
	}

	return channelIds, nil
}

func GetClosedChannelList() ([]model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	filter := bson.M{"channelStatus": model.CLOSED}
	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	var channels []model.Channel

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Println(err)
		}
		channels = append(channels, channel)
	}

	return channels, nil
}

func GetOpenedChannelList() ([]model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	filter := bson.M{"channelStatus": bson.M{
		"$not": bson.M{
			"$eq": model.CLOSED,
		},
	}}
	collection := database.Collection("channels")

	cur, err := collection.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}
	var channels []model.Channel

	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Println(err)
		}
		// To get the raw bson bytes use cursor.Current
		channels = append(channels, channel)
	}

	return channels, nil
}

func GetChannelsByChannelType(channelType model.ChannelType) ([]model.Channel, error){

	database, err := db.GetDatabase()
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"channelType": channelType,
	}

	collection := database.Collection("channels")
	cur, err := collection.Find(context.TODO(), filter)

	var channels []model.Channel
	defer cur.Close(context.Background())
	for cur.Next(context.Background()) {
		var channel model.Channel
		err := cur.Decode(&channel)
		if err != nil {
			log.Println(err)
		}
		// To get the raw bson bytes use cursor.Current
		channels = append(channels, channel)
	}

	return channels, nil
}

func GetChannelById(channelId int64) (model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return model.Channel{}, err
	}

	filter := bson.M{
		"channelId": channelId,
	}

	collection := database.Collection("channels")

	channel := model.Channel{}
	singleRecord := collection.FindOne(context.TODO(), filter)
	if err := singleRecord.Decode(&channel); err != nil {
		log.Println(err)
	}
	return channel, nil
}


func UpdateChannel(channel model.Channel) (model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return model.Channel{}, err
	}


	collection := database.Collection("channels")

	filter := bson.M{"channelId": channel.ChannelId}
	update := bson.M{"$set": bson.M{"channelStatus": channel.Status, "myBalance": channel.MyBalance, "otherAddress": channel.OtherAddress, "otherPort": channel.OtherPort, "otherIp": channel.OtherIp, "lockedBalance": channel.LockedBalance}}

	var rwMutex = new(sync.RWMutex)
	rwMutex.Lock()
	res, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println(err)
	}
	rwMutex.Unlock()

	fmt.Println(res.ModifiedCount)

	return channel, nil

}

func InsertChannel(channel model.Channel) (model.Channel, error) {

	database, err := db.GetDatabase()
	if err != nil {
		return model.Channel{}, err
	}

	collection := database.Collection("channels")

	insertResult, err := collection.InsertOne(context.TODO(), channel)
	if err != nil {
		return model.Channel{}, err
	}

	fmt.Println(insertResult.InsertedID)

	return channel, nil
}
