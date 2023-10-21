package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"urlshortener/internal/dto"
	"urlshortener/internal/utils"
)

type StorageInterface interface {
	PostLong(long string) (string, error)
	GetLong(short string) (string, error)
}

type server struct {
	g       *gin.Engine
	storage StorageInterface
}

func New(storage StorageInterface) *server {
	return &server{gin.Default(), storage}
}

func (s *server) Start(addr string) error {
	fmt.Println("Server running on port 8000")
	s.g.POST("/newUrl", func(context *gin.Context) {
		var request = new(dto.PostLongRequest)
		var response = new(dto.PostLongResponse)

		err := context.ShouldBindJSON(request)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = utils.ValidateUrl(request.Long); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("not a valid original URL: %w", err).Error()})
			return
		}

		response.Short, err = s.storage.PostLong(request.Long)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, response)
	})

	s.g.POST("/getLong", func(context *gin.Context) {
		var request = new(dto.GetLongRequest)
		var response = new(dto.GetLongResponse)

		err := context.ShouldBindJSON(request)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err = utils.ValidateUrl(request.Short); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("not a valid short URL: %w", err).Error()})
			return
		}

		response.Long, err = s.storage.GetLong(request.Short)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, response)
	})
	return s.g.Run(addr)
}
