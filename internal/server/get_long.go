package server

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"urlshortener/internal/utils"
	"urlshortener/pb"
)

func (s *server) GetLong(ctx context.Context, in *pb.GetLongRequest) (*pb.GetLongResponse, error) {
	shortUrl := in.GetShortUrl()

	if err := utils.ValidateUrl(shortUrl); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "GetLong: %v", err)
	}

	longUrl, err := s.storage.GetLong(shortUrl)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "GetLong: %v", err)
	}
	return &pb.GetLongResponse{LongUrl: longUrl}, nil
}
