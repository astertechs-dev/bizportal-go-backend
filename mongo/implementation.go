package mongo

import (
	"context"
	"errors"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoClient struct {
	cl *mongo.Client
}

type mongoDatabase struct {
	db *mongo.Database
}

type mongoCollection struct {
	coll *mongo.Collection
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

type mongoCursor struct {
	mc *mongo.Cursor
}

type mongoSession struct {
	mongo.Session
}

type nullAwareDecoder struct {
	defDecoder bsoncodec.ValueDecoder
	zeroValue  reflect.Value
}

func (d *nullAwareDecoder) DecodeValue(ctx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if err := vr.ReadNull(); err != nil {
		return d.defDecoder.DecodeValue(ctx, vr, val)
	}
	if !val.CanSet() {
		return errors.New("value not settable")
	}
	if err := vr.ReadNull(); err != nil {
		return err
	}

	val.Set(d.zeroValue)
	return nil
}

func NewClient(connection string) (Client, error) {
	time.Local = time.UTC
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connection))

	return &mongoClient{cl: c}, err
}

func (c *mongoClient) Database(name string) Database {
	return &mongoDatabase{db: c.cl.Database(name)}
}

func (c *mongoClient) Connect(ctx context.Context) error {
	return c.cl.Connect(ctx)
}

func (c *mongoClient) Disconnect(ctx context.Context) error {
	return c.cl.Disconnect(ctx)
}

func (c *mongoClient) StartSession() (mongo.Session, error) {
	session, err := c.cl.StartSession()
	return &mongoSession{session}, err
}

func (c *mongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return c.cl.UseSession(ctx, fn)
}

func (c *mongoClient) Ping(ctx context.Context) error {
	return c.cl.Ping(ctx, readpref.Primary())
}

func (d *mongoDatabase) Collection(name string) Collection {
	return &mongoCollection{coll: d.db.Collection(name)}
}

func (d *mongoDatabase) Client() Client {
	return &mongoClient{cl: d.db.Client()}
}

func (c *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	return &mongoSingleResult{sr: c.coll.FindOne(ctx, filter)}
}

func (c *mongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	mc, err := c.coll.Find(ctx, filter, opts...)
	return &mongoCursor{mc: mc}, err
}

func (c *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	return c.coll.InsertOne(ctx, document)
}

func (c *mongoCollection) InsertMany(ctx context.Context, document []interface{}) ([]interface{}, error) {
	res, err := c.coll.InsertMany(ctx, document)
	return res.InsertedIDs, err
}

func (c *mongoCollection) UpdateOne(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return c.coll.UpdateOne(ctx, filter, update, opts...)
}

func (c *mongoCollection) UpdateMany(
	ctx context.Context,
	filter interface{},
	update interface{},
	opts ...*options.UpdateOptions,
) (*mongo.UpdateResult, error) {
	return c.coll.UpdateMany(ctx, filter, update, opts...)
}

func (c *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	count, err := c.coll.DeleteOne(ctx, filter)
	return count.DeletedCount, err
}

func (c *mongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return c.coll.CountDocuments(ctx, filter, opts...)
}

func (c *mongoCollection) Aggregate(ctx context.Context, pipeline interface{}) (Cursor, error) {
	mc, err := c.coll.Aggregate(ctx, pipeline)
	return &mongoCursor{mc: mc}, err
}

func (sr *mongoSingleResult) Decode(v interface{}) error {
	return sr.sr.Decode(v)
}

func (mr *mongoCursor) Close(ctx context.Context) error {
	return mr.mc.Close(ctx)
}

func (mr *mongoCursor) Next(ctx context.Context) bool {
	return mr.mc.Next(ctx)
}

func (mr *mongoCursor) Decode(v interface{}) error {
	return mr.mc.Decode(v)
}

func (mr *mongoCursor) All(ctx context.Context, result interface{}) error {
	return mr.mc.All(ctx, result)
}
