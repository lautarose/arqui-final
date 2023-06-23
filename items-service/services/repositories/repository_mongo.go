package repositories

import (
	"context"
	"fmt"
	"items/dtos"
	model "items/models"
	e "items/utils/errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RepositoryMongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection string
}

func NewMongoDB(host string, port int, collection string) *RepositoryMongoDB {
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(fmt.Sprintf("mongodb://root:root@%s:%d/?authSource=admin&authMechanism=SCRAM-SHA-256", host, port)))
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	names, err := client.ListDatabaseNames(context.TODO(), bson.M{})
	if err != nil {
		panic(fmt.Sprintf("Error initializing MongoDB: %v", err))
	}

	fmt.Println("[MongoDB] Initialized connection")
	fmt.Println(fmt.Sprintf("[MongoDB] Available databases: %s", names))

	return &RepositoryMongoDB{
		Client:     client,
		Database:   client.Database("items"),
		Collection: collection,
	}
}

func (repo *RepositoryMongoDB) GetItemById(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return dtos.ItemDto{}, e.NewBadRequestApiError(fmt.Sprintf("error getting item %s invalid id", id))
	}
	result := repo.Database.Collection(repo.Collection).FindOne(context.TODO(), bson.M{
		"_id": objectID,
	})
	if result.Err() == mongo.ErrNoDocuments {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	var item model.Item
	if err := result.Decode(&item); err != nil {
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error getting item %s", id), err)
	}
	return dtos.ItemDto{
		Id:          id,
		Title:       item.Title,
		Seller:      item.Seller,
		Price:       item.Price,
		Currency:    item.Currency,
		Picture:     item.Picture,
		Description: item.Description,
		State:       item.State,
		City:        item.City,
		Street:      item.Street,
		Number:      item.Number,
	}, nil
}

func (repo *RepositoryMongoDB) InsertItems(ctx context.Context, items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError) {
	var itemsResponse dtos.ItemsDto
	for _, item := range items {
		result, err := repo.Database.Collection(repo.Collection).InsertOne(context.TODO(), model.Item{
			Title:       item.Title,
			Seller:      item.Seller,
			Price:       item.Price,
			Currency:    item.Currency,
			Picture:     item.Picture,
			Description: item.Description,
			State:       item.State,
			City:        item.City,
			Street:      item.Street,
			Number:      item.Number,
		})
		if err != nil {
			return dtos.ItemsDto{}, e.NewInternalServerApiError(fmt.Sprintf("error inserting item %s", item.Id), err)
		}
		item.Id = fmt.Sprintf(result.InsertedID.(primitive.ObjectID).Hex())

		itemsResponse = append(itemsResponse, item)
	}

	return itemsResponse, nil
}
