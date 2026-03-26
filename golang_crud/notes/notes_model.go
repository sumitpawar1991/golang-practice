package notes

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Notes struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Title     string             `bson:"title" json:"title"`
	Content   string             `bson:"content" json:"content"`
	pinned    bool               `bson:"pinned" json:"pinned"`
	createdAt time.Time          `bson:"createdAt" json:"createdAt"`
	updatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`
}

type CreateNoteRequest struct {
	Title   string `json:"title" binding:"requried"`
	Content string `json:"content" binding:"requried"`
	pinned  bool   `json:"pinned"`
}

type UpdateNoteRequest struct {
	Title   string `json:"title" binding:"requried"`
	Content string `json:"content" binding:"requried"`
	pinned  bool   `json:"pinned"`
}
