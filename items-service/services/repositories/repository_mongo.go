package repositories

import (
	"context"
	"fmt"
	"items/dtos"
	model "items/models"
	e "items/utils/errors/errors"
	"strconv"

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
	fmt.Printf("[MongoDB] Available databases: %s\n", names)

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
		UserID:      item.UserID,
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
			UserID:      item.UserID,
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
		item.Id = result.InsertedID.(primitive.ObjectID).Hex()

		itemsResponse = append(itemsResponse, item)
	}

	return itemsResponse, nil
}

func (repo *RepositoryMongoDB) UpdateItem(ctx context.Context, item dtos.ItemDto) (dtos.ItemDto, e.ApiError) {
	objectID, err := primitive.ObjectIDFromHex(item.Id)
	if err != nil {
		return dtos.ItemDto{}, e.NewBadRequestApiError("error updating item: invalid ID")
	}

	update := bson.M{
		"$set": bson.M{
			"userID":      item.UserID,
			"title":       item.Title,
			"seller":      item.Seller,
			"price":       item.Price,
			"currency":    item.Currency,
			"picture":     item.Picture,
			"description": item.Description,
			"state":       item.State,
			"city":        item.City,
			"street":      item.Street,
			"number":      item.Number,
		},
	}

	filter := bson.M{"_id": objectID}

	result, err := repo.Database.Collection(repo.Collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return dtos.ItemDto{}, e.NewInternalServerApiError(fmt.Sprintf("error updating item: %s", err.Error()), err)
	}

	if result.MatchedCount == 0 {
		return dtos.ItemDto{}, e.NewNotFoundApiError(fmt.Sprintf("item not found: %s", item.Id))
	}

	return item, nil
}

func (repo *RepositoryMongoDB) DeleteItem(ctx context.Context, id string) e.ApiError {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return e.NewBadRequestApiError(fmt.Sprintf("error deleting item %s: invalid id", id))
	}

	result, err := repo.Database.Collection(repo.Collection).DeleteOne(ctx, bson.M{"_id": objectID})
	if err != nil {
		return e.NewInternalServerApiError(fmt.Sprintf("error deleting item %s", id), err)
	}

	if result.DeletedCount == 0 {
		return e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
	}
	return nil
}

func (repo *RepositoryMongoDB) GetItemsIdByUserId(ctx context.Context, id string) ([]string, e.ApiError) {
	Id, err := strconv.Atoi(id)
	if err != nil {
		return nil, e.NewInternalServerApiError(fmt.Sprintf("error converting id to int: %s", err), err)
	}
	filter := bson.M{"user_id": Id} // Filtrar por el ID de usuario

	cursor, err := repo.Database.Collection(repo.Collection).Find(ctx, filter)
	if err != nil {
		return nil, e.NewInternalServerApiError(fmt.Sprintf("error getting items for user %s", id), err)
	}
	defer cursor.Close(ctx)

	var itemsIds []string

	for cursor.Next(ctx) {
		var item model.ItemEdit
		if err := cursor.Decode(&item); err != nil {
			return nil, e.NewInternalServerApiError(fmt.Sprintf("error decoding item for user %s", id), err)
		}

		itemsIds = append(itemsIds, item.Id.Hex())
	}

	if err := cursor.Err(); err != nil {
		return nil, e.NewInternalServerApiError(fmt.Sprintf("error iterating over items for user %s", id), err)
	}

	return itemsIds, nil
}
