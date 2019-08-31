package main

import (
	"context"
	"fmt"
	"log"
	"net"

	lis "github.com/thanhdvit/hr/db/lis"
	"github.com/thanhdvit/hr/proto"
	grpc "google.golang.org/grpc"
)

type EmployeeServiceServer struct {
}

func (s *EmployeeServiceServer) ReadBlog(ctx context.Context, req *proto.EmployeeReq) (*proto.EmployeeRes, error) {
	response := &proto.EmployeeRes{
		Employee: &proto.EmployeeReq{
			EmployeeId: 1,
		},
	}
	return response, nil
}

func main() {
	fmt.Println(lis.ValidUser("002099", "1"))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterEmployeeServiceServer(s, &EmployeeServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
