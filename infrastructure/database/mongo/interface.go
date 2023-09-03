package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	Collection(string) Collection
	Client() Client
}

type Collection interface {
	FindOne(context.Context, interface{}) SingleResult
	Find(context.Context, interface{}, ...*options.FindOptions) (Cursor, error)
	InsertOne(context.Context, interface{}) (interface{}, error)
	InsertMany(context.Context, []interface{}) ([]interface{}, error)
	UpdateOne(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	UpdateMany(context.Context, interface{}, interface{}, ...*options.UpdateOptions) (*mongo.UpdateResult, error)
	DeleteOne(context.Context, interface{}) (int64, error)
	CountDocuments(context.Context, interface{}, ...*options.CountOptions) (int64, error)
	Aggregate(context.Context, interface{}) (Cursor, error)
}

type SingleResult interface {
	Decode(interface{}) error
}

type Cursor interface {
	Decode(interface{}) error
	Next(context.Context) bool
	Close(context.Context) error
	All(context.Context, interface{}) error
}

type Client interface {
	Database(string) Database
	Connect(context.Context) error
	Disconnect(context.Context) error
	StartSession() (mongo.Session, error)
	UseSession(context.Context, func(mongo.SessionContext) error) error
	Ping(context.Context) error
}
