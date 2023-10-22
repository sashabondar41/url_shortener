package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"urlshortener/internal/utils"
	"urlshortener/pb"
)

func (s *server) ShortenLong(ctx context.Context, in *pb.ShortenLongRequest) (*pb.ShortenLongResponse, error) {
	longUrl := in.GetLongUrl()

	if err := utils.ValidateUrl(longUrl); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "ShortenLong: %v", err)
	}

	shortUrl, err := s.storage.PostLong(longUrl)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "ShortenLong: %v", err)
	}
	return &pb.ShortenLongResponse{ShortUrl: shortUrl}, nil
}
