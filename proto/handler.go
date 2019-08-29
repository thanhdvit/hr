package proto

import (
	context "context"

	hrpb "github.com/thanhdvit/hr/proto"
)

type EmployeeServiceServer struct {
}

func (s *EmployeeServiceServer) ReadBlog(ctx context.Context, req *hrpb.EmployeeReq) (*hrpb.EmployeeRes, error) {
	response := &hrpb.EmployeeRes{
		Employee: &hrpb.EmployeeReq{
			EmployeeId: 1,
		},
	}
	return response, nil
}
