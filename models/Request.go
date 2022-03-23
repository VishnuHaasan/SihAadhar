package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Request struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	Aadhaar    string             `json:"aadhaar" bson:"aadhaar"`
	RaisedTime int                `json:"raised_time" bson:"raised_time"`
	IsDone     bool               `json:"is_done" bson:"is_done"`
}

var requestCollection *mongo.Collection

func InitRequest(col *mongo.Collection) {
	requestCollection = col
}

func (r *Request) CreateRequest() (*Request, error) {
	_, err := requestCollection.InsertOne(context.TODO(), *r)
	return r, err
}

func UpdateRequest(aadhar string) (*Request, error) {
	after := options.After
	upsert := true
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	cursor := requestCollection.FindOneAndUpdate(context.Background(), bson.M{"aadhaar": aadhar, "is_done": false}, bson.M{"$set": bson.M{"is_done": true}}, &opt)
	updatedRequest := &Request{}
	if err := cursor.Decode(updatedRequest); err != nil {
		return nil, err
	}
	SendNotifications(aadhar)
	return updatedRequest, nil
}

func GetRequests() ([]Request, error) {
	cursor, err := requestCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	var res []Request
	if err = cursor.All(context.TODO(), &res); err != nil {
		return nil, err
	}
	return res, nil
}
