package main

import (
	"github.com/taha-ahmadi/iEvents/persistence"
	"github.com/taha-ahmadi/iEvents/persistence/mongodb"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYNAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {

	switch options {
	case MONGODB:
		return mongodb.NewMongoDBLayer(connection)
	}
	return nil, nil
}
