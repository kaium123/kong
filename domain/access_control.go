package domain

import (
	"context"
	"github.com/lib/pq"
)

type Api struct {
	Id       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string `json:"name"`
	Endpoint string `json:"endpoint"`
}

type Group struct {
	Id          int64         `json:"id" gorm:"primaryKey;autoIncrement;type:bigserial"`
	Name        string        `json:"name" gorm:"type:varchar(255);not null"`
	EndpointIds pq.Int64Array `json:"endpoint_ids" gorm:"type:bigint[]"`
	ApiKey      string        `json:"api_key" gorm:"type:varchar(255);not null"`
}

type ApiGroup struct {
	Id      int    `json:"id" gorm:"primaryKey;autoIncrement"`
	ApiId   string `json:"api_id"`
	GroupId string `json:"group_id"`
}

type ApiCriteria struct {
	Ids []int64
}

type GroupCriteria struct {
	ApiKey   string `json:"api_key"`
	Endpoint string `json:"endpoint"`
}

type AccessControlRepository interface {
	ApiStore(ctx context.Context, api *Api) (*Api, error)
	ApiGet(ctx context.Context, ctr *ApiCriteria) (*Api, error)
	ApiList(ctx context.Context, ctr *ApiCriteria) ([]*Api, error)
	ApiUpdate(ctx context.Context, api *Api) (*Api, error)
	ApiDelete(ctx context.Context, ctr *ApiCriteria) error
	GroupStore(ctx context.Context, group *Group) (*Group, error)
	GroupGet(ctx context.Context, ctr *GroupCriteria) (*Group, error)
	GroupList(ctx context.Context, ctr *GroupCriteria) ([]*Group, error)
	GroupUpdate(ctx context.Context, group *Group) (*Group, error)
	GroupDelete(ctx context.Context, ctr *GroupCriteria) error
}
type AccessControlService interface {
	ApiStore(ctx context.Context, api *Api) (*Api, error)
	ApiGet(ctx context.Context, ctr *ApiCriteria) (*Api, error)
	ApiList(ctx context.Context, ctr *ApiCriteria) ([]*Api, error)
	ApiUpdate(ctx context.Context, api *Api) (*Api, error)
	ApiDelete(ctx context.Context, ctr *ApiCriteria) error
	GroupStore(ctx context.Context, group *Group) (*Group, error)
	GroupGet(ctx context.Context, ctr *GroupCriteria) (*Group, error)
	GroupList(ctx context.Context, ctr *GroupCriteria) ([]*Group, error)
	GroupUpdate(ctx context.Context, group *Group) (*Group, error)
	GroupDelete(ctx context.Context, ctr *GroupCriteria) error
	Verify(ctx context.Context, api *GroupCriteria) (bool, error)
}
