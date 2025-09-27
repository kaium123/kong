package server

import (
	"context"
	"github.com/api_gaurd/domain"
	accesscontrolpb "github.com/api_gaurd/proto"
	"strconv"

	_ "gorm.io/gorm"
)

type AccessControlServer struct {
	service domain.AccessControlService
	accesscontrolpb.UnimplementedAccessControlServiceServer
}

func NewAccessControlServer(service domain.AccessControlService) AccessControlServer {
	return AccessControlServer{
		service: service,
	}
}

// ---- API CRUD ----
func (s *AccessControlServer) CreateApi(ctx context.Context, req *accesscontrolpb.CreateApiRequest) (*accesscontrolpb.CreateApiResponse, error) {
	api, err := s.service.ApiStore(ctx, &domain.Api{
		Name:     req.Name,
		Endpoint: req.Endpoint,
	})
	if err != nil {
		return nil, err
	}

	return &accesscontrolpb.CreateApiResponse{Api: &accesscontrolpb.Api{
		Name:     api.Name,
		Id:       strconv.Itoa(api.Id),
		Endpoint: api.Endpoint,
	}}, nil
}

func (s *AccessControlServer) GetApi(ctx context.Context, req *accesscontrolpb.GetApiRequest) (*accesscontrolpb.GetApiResponse, error) {
	api, err := s.service.ApiGet(ctx, &domain.ApiCriteria{})
	if err != nil {
		return nil, err
	}
	return &accesscontrolpb.GetApiResponse{
		Api: &accesscontrolpb.Api{
			Name:     api.Name,
			Id:       strconv.Itoa(api.Id),
			Endpoint: api.Endpoint,
		},
	}, nil

}

//
//func (s *AccessControlServer) UpdateApi(ctx context.Context, req *accesscontrolpb.UpdateApiRequest) (*accesscontrolpb.UpdateApiResponse, error) {
//	api, err := s.service.ApiUpdate(ctx, &domain.Api{
//		Id:       req.Id,
//		Name:     req.Name,
//		Endpoint: req.Endpoint,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return &accesscontrolpb.UpdateApiResponse{
//		Api: &accesscontrolpb.Api{
//			Name:     api.Name,
//			Id:       api.Id,
//			Endpoint: api.Endpoint,
//		},
//	}, nil
//}
//
//func (s *AccessControlServer) DeleteApi(ctx context.Context, req *accesscontrolpb.DeleteApiRequest) (*accesscontrolpb.DeleteApiResponse, error) {
//	err := s.service.ApiDelete(ctx, &domain.ApiCriteria{})
//	if err != nil {
//		return nil, err
//	}
//	return &accesscontrolpb.DeleteApiResponse{
//		Success: true,
//	}, nil
//}
//
//func (s *AccessControlServer) ListApis(ctx context.Context, req *accesscontrolpb.ListApisRequest) (*accesscontrolpb.ListApisResponse, error) {
//	_, err := s.service.ApiList(ctx, &domain.ApiCriteria{})
//	if err != nil {
//		return nil, err
//	}
//	return &accesscontrolpb.ListApisResponse{
//
//	}, nil
//}

// ---- GROUP CRUD ----
func (s *AccessControlServer) CreateGroup(ctx context.Context, req *accesscontrolpb.CreateGroupRequest) (*accesscontrolpb.CreateGroupResponse, error) {

	apiIds := []int64{}
	for _, endpoint := range req.Endpoints {
		api, err := s.service.ApiStore(ctx, &domain.Api{
			Endpoint: endpoint,
		})
		if err != nil {
			return nil, err
		}

		apiIds = append(apiIds, int64(api.Id))
	}

	group, err := s.service.GroupStore(ctx, &domain.Group{
		Name:        req.Name,
		EndpointIds: apiIds,
		ApiKey:      req.ApiKey,
	})
	if err != nil {
		return nil, err
	}

	return &accesscontrolpb.CreateGroupResponse{
		Group: &accesscontrolpb.Group{
			Name:   group.Name,
			Id:     group.Id,
			ApiKey: group.ApiKey,
		},
	}, nil

}

func (s *AccessControlServer) Verify(ctx context.Context, req *accesscontrolpb.VerifyRequest) (*accesscontrolpb.VerifyResponse, error) {
	isVerified, err := s.service.Verify(ctx, &domain.GroupCriteria{
		ApiKey:   req.ApiKey,
		Endpoint: req.Endpoint,
	})
	if err != nil {
		return nil, err
	}

	return &accesscontrolpb.VerifyResponse{
		Verified: isVerified,
	}, nil
}
