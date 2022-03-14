package db

import (
	"context"
	"errors"
	"fmt"

	"users/internal/serviceerror"
	"users/internal/users"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type db struct {
	collection *mongo.Collection
}

func NewStorage(database *mongo.Database, collection string) users.Storage {
	database.Collection(collection)
	return &db{
		collection: database.Collection(collection),
	}
}

func (db *db) Create(ctx context.Context, user users.User) (str string, errRes *serviceerror.ErrorResponse) {
	result, err := db.collection.InsertOne(ctx, user)
	if err != nil {
		err = fmt.Errorf("failed to create user due to error: %v", err)
		return str, serviceerror.NewServiceError(err, "", "", "")
	}
	oid, ok := result.InsertedID.(primitive.ObjectID)
	if ok {
		return oid.Hex(), nil
	}
	err = fmt.Errorf("failed to convert objectid to hex. probably oid: %s", oid)
	return "", serviceerror.NewServiceError(err, "", "", "")
}

func (db *db) FindAll(ctx context.Context) (user []users.User, errRes *serviceerror.ErrorResponse) {
	cursor, err := db.collection.Find(ctx, bson.M{})
	if cursor.Err() != nil {
		err = fmt.Errorf("failed to find all users due to error: %v", err)
		errRes = serviceerror.NewServiceError(err, "", "", "")
		return user, errRes
	}

	if err = cursor.All(ctx, &user); err != nil {
		err = fmt.Errorf("failed to read all documents from cursor. error: %v", err)
		errRes = serviceerror.NewServiceError(err, "", "", "")
		return user, errRes
	}
	return user, nil
}

func (db *db) FindOne(ctx context.Context, id string) (user users.User, errRes *serviceerror.ErrorResponse) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err = fmt.Errorf("failed to convert hex to objectid. hex: %s", id)
		errRes = serviceerror.NewServiceError(err, "", "", "")
		return user, errRes
	}
	filter := bson.M{"_id": oid}
	result := db.collection.FindOne(ctx, filter)
	if result.Err() != nil {
		if errors.Is(result.Err(), mongo.ErrNoDocuments) {
			return user, serviceerror.ErrorNotFound
		}
		err = fmt.Errorf("failed to find one user by id: %s due to error: %v", id, err)
		errRes = serviceerror.NewServiceError(err, "", "", "")
		return user, errRes
	}
	if err = result.Decode(&user); err != nil {
		err = fmt.Errorf("failed to decode user (id:%s) from DB due to error: %v", id, err)
		errRes = serviceerror.NewServiceError(err, "", "", "")
		return user, errRes
	}
	return user, errRes
}

func (db *db) Update(ctx context.Context, user users.User) *serviceerror.ErrorResponse {
	objectID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		err = fmt.Errorf("failed to convert user ID to ObjectID. ID=%s", user.ID)
		return serviceerror.NewServiceError(err, "", "", "")
	}

	filter := bson.M{"_id": objectID}

	userBytes, err := bson.Marshal(user)
	if err != nil {
		err = fmt.Errorf("failed to marshal user. error: %v", err)
		return serviceerror.NewServiceError(err, "", "", "")
	}

	var updateUserObj bson.M
	err = bson.Unmarshal(userBytes, &updateUserObj)
	if err != nil {
		err = fmt.Errorf("failed to unmarshal user bytes. error: %v", err)
		return serviceerror.NewServiceError(err, "", "", "")
	}

	delete(updateUserObj, "_id")

	update := bson.M{
		"$set": updateUserObj,
	}

	result, err := db.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		err = fmt.Errorf("failed to execute update user query. error: %v", err)
		return serviceerror.NewServiceError(err, "", "", "")
	}

	if result.MatchedCount == 0 {
		return serviceerror.ErrorNotFound
	}

	return nil

}

func (db *db) Delete(ctx context.Context, id string) *serviceerror.ErrorResponse {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		err = fmt.Errorf("failed to convert user ID to ObjectID. ID=%s", id)
		return serviceerror.NewServiceError(err, "", "", "")
	}

	filter := bson.M{"_id": objectID}

	result, err := db.collection.DeleteOne(ctx, filter)
	if err != nil {
		err = fmt.Errorf("failed to execute query. error: %v", err)
		return serviceerror.NewServiceError(err, "", "", "")
	}
	if result.DeletedCount == 0 {
		return serviceerror.ErrorNotFound
	}
	return nil
}
