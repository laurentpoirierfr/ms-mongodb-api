package repositories

import (
	"context"
	"os"
	"reflect"
	"time"

	"github.com/gookit/slog"
	"github.com/laurentpoirierfr/ms-mongodb-api/internal/core/domain"
	"github.com/laurentpoirierfr/ms-mongodb-api/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoRepository struct {
	db          *mongo.Database
	collections map[string]*mongo.Collection
}

func NewMongoRepository(config *util.Config) *MongoRepository {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.DatabaseUrl).SetServerAPIOptions(serverAPI)
	// Declare Context type object for managing multiple API requests
	ctx, _ := context.WithTimeout(context.TODO(), time.Duration(config.ConnectionTimeout))
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		slog.FatalErr(err)
		os.Exit(-1)
	}
	slog.Info("MongoDb connection ok.")
	return &MongoRepository{
		db:          client.Database(config.DatabaseName),
		collections: make(map[string]*mongo.Collection),
	}
}

func (repo *MongoRepository) GetDocuments(ctx context.Context, documents string, offset, limit int64) (interface{}, error) {
	coll := repo.getCollection(documents)

	findOptions := options.Find()
	findOptions.SetSkip(offset)
	findOptions.SetLimit(limit)

	cursor, err := coll.Find(ctx, bson.D{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	emptyValue := repo.getEmptyValue()
	results := reflect.New(reflect.SliceOf(reflect.TypeOf(emptyValue))).Interface()

	if err := cursor.All(ctx, results); err != nil {
		return nil, err
	}

	totalCount, _ := coll.CountDocuments(ctx, bson.M{})

	paginatedResponse := domain.PaginatedResponse{
		Offset:    offset,
		Limit:     limit,
		Total:     totalCount,
		Documents: results,
	}

	return paginatedResponse, nil
}

func (repo *MongoRepository) GetDocumentById(ctx context.Context, documents string, id string) (interface{}, error) {
	coll := repo.getCollection(documents)
	emptyValue := repo.getEmptyValue()
	result := reflect.New(reflect.TypeOf(emptyValue)).Interface()
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = coll.FindOne(ctx, bson.M{"_id": objectID}).Decode(result)
	if err != nil {
		// if err == mongo.ErrNoDocuments {
		// 	return nil, err
		// }
		return nil, err
	}

	return result, nil
}

func (repo *MongoRepository) CreateDocument(ctx context.Context, documents string, document interface{}) (interface{}, error) {
	coll := repo.getCollection(documents)
	result, err := coll.InsertOne(ctx, document)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *MongoRepository) UpdateDocument(ctx context.Context, documents string, document interface{}, id string) (interface{}, error) {
	coll := repo.getCollection(documents)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	updateResult, err := coll.UpdateOne(
		ctx,
		bson.M{"_id": objectID},
		bson.D{{Key: "$set", Value: document}},
	)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func (repo *MongoRepository) DeleteDocument(ctx context.Context, documents string, id string) (interface{}, error) {
	coll := repo.getCollection(documents)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	deleteResult, err := coll.DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return nil, err
	}
	return deleteResult, nil
}

func (repo *MongoRepository) getCollection(documents string) *mongo.Collection {
	if repo.collections[documents] == nil {
		repo.collections[documents] = repo.db.Collection(documents)
	}
	collection := repo.collections[documents]
	return collection
}

func (repo *MongoRepository) getEmptyValue() interface{} {
	return &map[string]any{}
}
