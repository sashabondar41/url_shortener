package server

import (
	"urlshortener/pb"
)

type StorageInterface interface {
	PostLong(long string) (string, error)
	GetLong(short string) (string, error)
}

type server struct {
	pb.UnimplementedUrlShortenerServer
	storage StorageInterface
}

func New(storage StorageInterface) *server {
	return &server{storage: storage}
}
