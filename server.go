package main

import (
	"context"
	"fmt"
	"log"
	"net"

	lis "github.com/thanhdvit/hr/db/lis"
	hrpd "github.com/thanhdvit/hr/hello"
	grpc "google.golang.org/grpc"
)

type EmployeeServiceServer struct {
}

func (s *EmployeeServiceServer) ReadBlog(ctx context.Context, req *hrpd.EmployeeReq) (*hrpd.EmployeeRes, error) {
	response := &hrpd.EmployeeRes{
		Employee: &hrpd.EmployeeReq{
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
	hrpd.RegisterServiceServer(s, &EmployeeServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
