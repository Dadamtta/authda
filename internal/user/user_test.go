package user

import (
	"context"
	"dadamtta/private/p_appl"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Restaurant struct {
	Name         string
	RestaurantId string        `bson:"restaurant_id,omitempty"`
	Cuisine      string        `bson:"cuisine,omitempty"`
	Address      interface{}   `bson:"address,omitempty"`
	Borough      string        `bson:"borough,omitempty"`
	Grades       []interface{} `bson:"grades,omitempty"`
}

func TestInsertNoSQLWeddingInvitation(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	weddingInvitation := p_appl.WeddingInvitation{Id: "3122334456", BrideFirst: false}
	coll := client.Database("ddt_test").Collection("apps")

	_, err = coll.InsertOne(context.TODO(), weddingInvitation)
	if err != nil {
		panic(err)
	}
}

func TestReplaceNoSQLWeddingInvitation(t *testing.T) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	weddingInvitation := p_appl.WeddingInvitation{Id: "1122334455", BrideFirst: false}
	filter := bson.D{{"_id", "1122334455"}}
	// newRestaurant := Restaurant{Name: "8282", Cuisine: "Korean"}
	coll := client.Database("ddt_test").Collection("apps")

	_, err = coll.ReplaceOne(context.TODO(), filter, weddingInvitation)

	if err != nil {
		panic(err)
	}
}
