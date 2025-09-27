package service

import (
	"context"
	"github.com/api_gaurd/domain"
	_ "gorm.io/gorm"
)

type AccessControlService struct {
	repo domain.AccessControlRepository
}

func NewAccessControlService(repo domain.AccessControlRepository) domain.AccessControlService {
	return &AccessControlService{
		repo: repo,
	}
}

func (r *AccessControlService) ApiStore(ctx context.Context, api *domain.Api) (*domain.Api, error) {
	return r.repo.ApiStore(ctx, api)
}

func (r *AccessControlService) ApiGet(ctx context.Context, ctr *domain.ApiCriteria) (*domain.Api, error) {
	return r.repo.ApiGet(ctx, ctr)
}

func (r *AccessControlService) ApiList(ctx context.Context, ctr *domain.ApiCriteria) ([]*domain.Api, error) {
	return r.repo.ApiList(ctx, ctr)
}

func (r *AccessControlService) ApiUpdate(ctx context.Context, address *domain.Api) (*domain.Api, error) {

	return r.repo.ApiUpdate(ctx, address)
}

func (r *AccessControlService) ApiDelete(ctx context.Context, ctr *domain.ApiCriteria) error {

	return r.repo.ApiDelete(ctx, ctr)
}

func (r *AccessControlService) GroupStore(ctx context.Context, api *domain.Group) (*domain.Group, error) {
	return r.repo.GroupStore(ctx, api)
}

func (r *AccessControlService) GroupGet(ctx context.Context, ctr *domain.GroupCriteria) (*domain.Group, error) {
	return r.repo.GroupGet(ctx, ctr)
}

func (r *AccessControlService) GroupList(ctx context.Context, ctr *domain.GroupCriteria) ([]*domain.Group, error) {
	return r.repo.GroupList(ctx, ctr)
}

func (r *AccessControlService) GroupUpdate(ctx context.Context, group *domain.Group) (*domain.Group, error) {
	return r.repo.GroupUpdate(ctx, group)
}

func (r *AccessControlService) GroupDelete(ctx context.Context, ctr *domain.GroupCriteria) error {
	return r.repo.GroupDelete(ctx, ctr)
}

func (r *AccessControlService) Verify(ctx context.Context, ctr *domain.GroupCriteria) (bool, error) {
	group, err := r.GroupGet(ctx, ctr)
	if err != nil {
		return false, err
	}

	apiList, err := r.ApiList(ctx, &domain.ApiCriteria{
		Ids: group.EndpointIds,
	})

	for _, api := range apiList {
		if ctr.Endpoint == api.Endpoint {
			return true, nil
		}
	}

	return false, nil
}
