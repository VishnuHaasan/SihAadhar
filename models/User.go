package models

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceProvider struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id"`
	Name     string             `json:"name" bson:"name"`
	Url      string             `json:"url" bson:"url"`
	IsActive bool               `json:"is_active" bson:"is_active"`
}

type User struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	AadhaarID        string             `json:"aadhaar" bson:"aadhaar"`
	ServiceProviders []ServiceProvider  `json:"service_providers" bson:"service_providers"`
}

var userCollection *mongo.Collection

func InitUser(col *mongo.Collection) {
	userCollection = col
}

func (user *User) CreateUser() (*User, error) {
	_, err := userCollection.InsertOne(context.TODO(), *user)
	return user, err
}

func GetUser(AadhaarID string) (*User, error) {
	res := &User{}
	err := userCollection.FindOne(context.TODO(), bson.M{"aadhaar": AadhaarID}).Decode(res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func AddServiceProvider(sp *ServiceProvider, Aadhaar string) (*User, error) {
	after := options.After
	upsert := true
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	cursor := userCollection.FindOneAndUpdate(context.Background(), bson.M{"aadhaar": Aadhaar}, bson.M{"$push": bson.M{"service_providers": *sp}}, &opt)
	updatedUser := &User{}
	if err := cursor.Decode(updatedUser); err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func ToggleIsActiveForProvider(Aadhaar string, Id primitive.ObjectID, IsActive bool) (*User, error) {
	after := options.After
	upsert := true
	opt := options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}
	cursor := userCollection.FindOneAndUpdate(context.Background(), bson.M{"aadhaar": Aadhaar, "service_providers._id": Id}, bson.M{"$set": bson.M{"service_providers.$.is_active": IsActive}}, &opt)
	updatedUser := &User{}
	if err := cursor.Decode(updatedUser); err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func SendNotifications(aadhaar string) error {
	user, err := GetUser(aadhaar)
	if err != nil {
		return err
	}
	for i := 0; i < len(user.ServiceProviders); i++ {
		if user.ServiceProviders[i].IsActive {
			go SendRequest(user.ServiceProviders[i].Url)
		}
	}
	return nil
}

func SendRequest(url string) {
	http.Get(url)
}
