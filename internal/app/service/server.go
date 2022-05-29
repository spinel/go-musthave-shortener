package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/spinel/go-musthave-shortener/internal/app/model"
	"github.com/spinel/go-musthave-shortener/internal/app/pkg/pb"
	"github.com/spinel/go-musthave-shortener/internal/app/repository"
)

type Server struct {
	Repo repository.URLStorer
}

func (s *Server) GetURL(ctx context.Context, req *pb.GetURLRequest) (*pb.GetURLResponse, error) {
	e, err := s.Repo.GetURL(ctx, req.UrlCode)
	if err != nil {
		return &pb.GetURLResponse{
			Error: "user already exists",
		}, nil
	}

	return &pb.GetURLResponse{
		EntityId: int64(e.ID),
	}, nil
}

func (s *Server) CreateURL(ctx context.Context, req *pb.CreateURLRequest) (*pb.CreateURLResponse, error) {
	e := &model.Entity{
		Code: req.Code,
		URL:  req.Url,
	}
	e, err := s.Repo.CreateURL(ctx, e)
	if err != nil {
		return &pb.CreateURLResponse{
			Error: "user already exists",
		}, nil
	}

	return &pb.CreateURLResponse{
		EntityId: int64(e.ID),
	}, nil
}

func (s *Server) EnqueueDelete(ctx context.Context, req *pb.EnqueueDeleteRequest) (*pb.EnqueueDeleteResponse, error) {
	uuid, _ := uuid.Parse(req.UserUuid)
	err := s.Repo.EnqueueDelete(req.Codes, uuid)
	return &pb.EnqueueDeleteResponse{
		Error: err.Error(),
	}, nil
}
