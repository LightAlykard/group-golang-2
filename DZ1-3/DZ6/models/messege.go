package models

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	ID primitive.ObjectID `json:"id"            bson:"_id,omitempty" `
}

type Messege struct {
	Mongo `inline`
	Name  string `json:"Name"      bson:"Name"`
	Text  string `json:"Text"       bson:"Text"`
}
type Messeges map[string]Messege

//Get Posts
func Get(ctx context.Context, db *mongo.Database) (*Messeges, error) {
	cur, err := db.Collection("messeges").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	messeges := make(Messeges, 10)

	defer cur.Close(ctx)
	for cur.Next(ctx) {
		messege := Messege{}
		err := cur.Decode(&messege)
		if err != nil {
			return nil, err
		}
		messeges[messege.ID.Hex()] = messege
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return &messeges, nil
}
func (messeges *Messeges) NewPost(ctx context.Context, db *mongo.Database, name string, text string) (string, error) {
	tmp := Messege{
		Name: Name,
		Text: text,
	}

	done, err := db.Collection("messeges").InsertOne(ctx, tmp)
	fmt.Println(done)
	if err != nil {
		return "", err
	}
	id := done.InsertedID.(primitive.ObjectID).Hex()
	(*messeges)[id] = tmp
	return id, nil
}

//UpdatePost
func (messeges *messeges) UpdatePost(ctx context.Context, db *mongo.Database, id string, name string, text string) error {
	t := time.Now()
	tmp := Messege{
		Name: name,
		Text: text,
	}

	newID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	tmp.ID = newID

	rslt, err := db.Collection("messeges").ReplaceOne(ctx, bson.M{"_id": newID}, tmp)
	fmt.Println(rslt)
	if err != nil {
		return err
	}
	(*messeges)[id] = tmp
	return nil
}

//DeletePost
func (messeges *Messeges) DeletePost(ctx context.Context, db *mongo.Database, id string) error {
	messegeID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	rslt, err := db.Collection("messeges").DeleteOne(ctx, bson.M{"_id": messegeID})
	fmt.Println(rslt)
	if err != nil {
		return err
	}
	delete((*messeges), id)
	return nil
}
