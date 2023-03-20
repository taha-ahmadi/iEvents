package mongodb

import (
	"context"

	"github.com/pkg/errors"
	"github.com/taha-ahmadi/iEvents/persistence"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DB     = "myevents"
	USERS  = "users"
	EVENTS = "events"
)

type MongoDBLayer struct {
	client *mongo.Client
}

func NewMongoDBLayer(connection string) (persistence.DatabaseHandler, error) {
	// Set up MongoDB client
	clientOptions := options.Client().ApplyURI(connection)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	// Verify connection
	err = client.Ping(context.Background(), readpref.Primary())
	if err != nil {
		return nil, errors.Wrap(err, "failed to ping database")
	}
	return &MongoDBLayer{client: client}, nil
}

func (mgoLayer *MongoDBLayer) AddEvent(e persistence.Event) ([]byte, error) {
	collection := mgoLayer.client.Database(DB).Collection(EVENTS)

	result, err := collection.InsertOne(context.Background(), e)
	if err != nil {
		return nil, errors.Wrap(err, "failed to insert event")
	}

	return []byte(result.InsertedID.(primitive.ObjectID).Hex()), nil
}

func (mgoLayer *MongoDBLayer) FindEvent(id []byte) (persistence.Event, error) {
	collection := mgoLayer.client.Database(DB).Collection(EVENTS)

	objectID, err := primitive.ObjectIDFromHex(string(id))
	if err != nil {
		return persistence.Event{}, errors.Wrap(err, "invalid event ID")
	}

	filter := bson.M{"_id": objectID}

	var event persistence.Event
	err = collection.FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return persistence.Event{}, errors.Wrap(err, "event not found")
		}
		return persistence.Event{}, errors.Wrap(err, "failed to find event")
	}

	return event, nil
}

func (mgoLayer *MongoDBLayer) FindEventByName(name string) (persistence.Event, error) {
	collection := mgoLayer.client.Database(DB).Collection(EVENTS)

	filter := bson.M{"name": name}

	var event persistence.Event
	err := collection.FindOne(context.Background(), filter).Decode(&event)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return persistence.Event{}, errors.Wrap(err, "event not found")
		}
		return persistence.Event{}, errors.Wrap(err, "failed to find event")
	}

	return event, nil
}

func (mgoLayer *MongoDBLayer) FindAllAvailableEvents() ([]persistence.Event, error) {
	collection := mgoLayer.client.Database(DB).Collection(EVENTS)

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to find events")
	}

	events := []persistence.Event{}
	if err = cursor.All(context.Background(), &events); err != nil {
		return nil, errors.Wrap(err, "failed to parse events")
	}

	return events, nil
}
