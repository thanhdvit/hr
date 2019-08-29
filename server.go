package main

import (
	"fmt"
	"log"
	"net"

	lis "github.com/thanhdvit/hr/db/lis"
	pd "github.com/thanhdvit/hr/impl"
	grpc "google.golang.org/grpc"
)

func main() {
	fmt.Println(lis.ValidUser("002099", "1"))

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pd.RegisterServiceServer(s, &EmployeeServiceServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
