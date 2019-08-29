package server

import (
	context "context"

	hrpd "github.com/thanhdvit/hr/proto"
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
