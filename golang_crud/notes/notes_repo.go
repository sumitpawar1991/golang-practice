package notes

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

//it data access layer

type Repo struct {
	coll *mongo.Collection
}

func NewRepo(db *mongo.Database) *Repo {
	return &Repo{
		coll: db.Collection("notes"),
	}
}

func (r *Repo) Create(ctx context.Context, note Note) (Note, error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	_, err := r.coll.InsertOne(opCtx, note)

	if err != nil {
		return Note{}, fmt.Errorf("Insert note failed")
	}

	return note, nil
}

func (r *Repo) List(ctx context.Context) ([]Note, error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{} //match all the docs

	//matching elements
	cursor, err := r.coll.Find(opCtx, filter)

	if err != nil {
		return nil, fmt.Errorf("find notes failed: %w", err)
	}

	//cursor must be closed after use // avoid any kinds of leaks
	defer cursor.Close(opCtx)

	var notes []Note

	if err := cursor.All(opCtx, &notes); err != nil {
		return nil, fmt.Errorf("Decode notes failed:%w", err)
	}

	return notes, nil

}

func (r *Repo) GetByID(ctx context.Context, id primitive.ObjectID) (Note, error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{"_id": id}

	var note Note

	err := r.coll.FindOne(opCtx, filter, options.FindOne()).Decode(&note)

	if err != nil {
		returm Note {}, fmt.Errorf("Find note by failed: %w",err)
	}

	return note , nil
}

func (r *Repo ) UpdateByID(ctx context.Context, id primitive.ObjectID,req UpdateNoteRequest) {
	opctx , cancel := context.WithTimeout(ctx , 5*time.Second)

	defer cancel()

	filter := bson.M{"_id":id}

	update := bson.M{
		"$set" : bson.M{
			"title" : req.Title,
			"content" : req.Content,
			"pinned" : req.pinned,
			"UpdatedAt" : time.Now().UTC(),

		}
	}


	after := options.After

	opts := options.FindOneAndUpdateOptions {
		ReturnDocument: &after,
	}

	var updated Note
	err := r.coll.FindOneAndUpdate(opctx,filter,update , &opts).Decode(&updated)

	if err != nil {
		return Note{},fmt.Errorf("Update note failed: %w",err)
	}

	return updated , nil
}


func (r *Repo) DeleteByID(ctx context.Context , id Primitive.ObjectID) (bool,error) {

	opCtx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	filter := bson.M{"_id":id}

	res.err := r.coll.DeleteOne(opCtx, filter)

	if err != nil {
		return false , fmt.Errorf("failed to delete the given note :%w",err)
	}

	if (res.DeletedCount===0) {
		return false,nil
	}

	return true,nil

}