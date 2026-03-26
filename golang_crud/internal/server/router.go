package server

import (
	"golang-crud/notes"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func NewRouter(database *mongo.Database) *gin.Engine {

	//default router

	r := gin.Default()

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":     true,
			"status": "healtthy",
		})
	})

	notes.RegisterRoutes(r, database)

	return r
}
