package service

import (
	"context"

	pb "github.com/jgilliescommure/go/api/personcrud"
)

type PersoncrudService struct {
	pb.UnimplementedPersoncrudServer
}

func NewPersoncrudService() *PersoncrudService {
	return &PersoncrudService{}
}

func (s *PersoncrudService) CreatePerson(ctx context.Context, req *pb.CreatePersonRequest) (*pb.CreatePersonReply, error) {
	return &pb.CreatePersonReply{Id: 11}, nil
}
func (s *PersoncrudService) UpdatePerson(ctx context.Context, req *pb.UpdatePersonRequest) (*pb.PersonReply, error) {
	return &pb.PersonReply{}, nil
}
func (s *PersoncrudService) GetPerson(ctx context.Context, req *pb.GetPersonRequest) (*pb.PersonReply, error) {
	return &pb.PersonReply{}, nil
}
func (s *PersoncrudService) ListPersons(ctx context.Context, req *pb.ListPersonRequest) (*pb.ListPersonReply, error) {
	return &pb.ListPersonReply{}, nil
}
