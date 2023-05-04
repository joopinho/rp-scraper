package grps

import (
	"context"
	"log"

	pb "github.com/joopinho/rp-scarper/api/proto/v1/organization"
	"github.com/joopinho/rp-scarper/internal/services"
	"github.com/joopinho/rp-scarper/internal/tools"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrganizationGrpcServer struct {
	Svc *services.OrganizationService
	pb.UnimplementedOrganizationServiceServer
}

func (s *OrganizationGrpcServer) QueryProfile(c context.Context, request *pb.QueryRequest) (*pb.ProfileResponse, error) {

	p, err := s.Svc.GetOrganization(request.GetInn())
	if err != nil {
		log.Printf("[inn: %s] - %v", request.GetInn(), err.Error())
		return nil, status.Errorf(codes.Code(err.(*tools.ServiceError).Code), err.Error())
	}

	return &pb.ProfileResponse{Inn: p.INN(), Kpp: p.KPP(), Ceo: p.CEO(), Name: p.Name()},
		status.Errorf(codes.OK, "")
}
