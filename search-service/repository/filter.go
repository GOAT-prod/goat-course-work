package repository

import (
	"github.com/GOAT-prod/goatcontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"search-service/database"
)

type Filter interface {
	GetFilters(ctx goatcontext.Context) (filters []database.Filter, err error)
}

type FilterRepository struct {
	mongo      *mongo.Client
	database   string
	collection string
}

func NewFilterRepository(mongo *mongo.Client, database, collection string) Filter {
	return &FilterRepository{
		mongo:      mongo,
		database:   database,
		collection: collection,
	}
}

func (r *FilterRepository) GetFilters(ctx goatcontext.Context) (filters []database.Filter, err error) {
	collection := r.mongo.Database(r.database).Collection(r.collection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	return filters, cursor.All(ctx, &filters)
}
