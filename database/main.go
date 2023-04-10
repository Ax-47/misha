package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	Colls  map[string]*mongo.Collection
}

func Get() *Database {
	db := &Database{}
	return db
}
func (db *Database) Init(uri, database string, colls []string) error {
	var err error
	db.Colls = make(map[string]*mongo.Collection)
	db.Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	for _, coll := range colls {
		db.Colls[coll] = db.Client.Database(database).Collection(coll)
	}

	return err
}
