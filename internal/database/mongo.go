package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDriver struct {
	host string
	port string
	user string
	pass string
	name string
}

func NewMongoDriver(host string, port string, user string, pass string, name string) *MongoDriver {
	return &MongoDriver{host: host, port: port, user: user, pass: pass, name: name}
}

func (d MongoDriver) GetClient() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(20)*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%s", d.user, d.pass, d.host, d.port)))

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (d MongoDriver) GetDatabase() (*mongo.Database, error) {
	client, err := d.GetClient()
	if err != nil {
		return nil, err
	}

	return client.Database(d.name), nil
}
