package repository

import (
	"context"
	"fmt"

	"github.com/api_gaurd/domain"
	"github.com/api_gaurd/pkg/conn"
	_ "gorm.io/gorm"
)

type AccessControlPostgreSQL struct {
	db *conn.DB
}

func NewAccessControlPostgreSQL(db *conn.DB) domain.AccessControlRepository {
	return &AccessControlPostgreSQL{
		db: db,
	}
}

func (r *AccessControlPostgreSQL) ApiStore(ctx context.Context, api *domain.Api) (*domain.Api, error) {
	q := r.db.DB

	err := q.WithContext(ctx).Create(api).Error
	if err != nil {
		return nil, fmt.Errorf("repository:postgreSQL: failed to Store customer: %v", err)
	}

	return api, nil
}

func (r *AccessControlPostgreSQL) ApiGet(ctx context.Context, ctr *domain.ApiCriteria) (*domain.Api, error) {
	q := r.db.DB

	qry := q.WithContext(ctx)

	//if ctr.Id != nil && *ctr.Id > 0 {
	//	qry = qry.Where("id = ?", ctr.Id)
	//}
	//
	//if ctr.Phone != nil {
	//	qry = qry.Where("phone = ?", ctr.Phone)
	//}

	var api domain.Api
	if err := qry.First(&api).Error; err != nil {
		return nil, fmt.Errorf("repository:postgreSQL: failed to FetchOne customer: %v", err)
	}

	return &api, nil
}

func (r *AccessControlPostgreSQL) ApiList(ctx context.Context, ctr *domain.ApiCriteria) ([]*domain.Api, error) {

	qry := r.db.DB.WithContext(ctx) // operation on gorm db

	if len(ctr.Ids) > 0 {
		qry = qry.Where("id IN(?)", ctr.Ids)
	}

	var apiList []*domain.Api

	if err := qry.Find(&apiList).Error; err != nil {
		return nil, fmt.Errorf("repository:postgreSQL: failed to Fetch customer: %v", err)
	}

	return apiList, nil
}

func (r *AccessControlPostgreSQL) ApiUpdate(ctx context.Context, address *domain.Api) (*domain.Api, error) {

	/*
	   if customer.Id == "" {
	       return nil, errors.New("repository:mongodb: Update failed: customer id required")
	   }

	   if err := r.db.DB.WithContext(ctx).Unscoped().Updates(customer).Error; err != nil {
	       return fmt.Errorf("repository:postgreSQLdb: Update failed: %v", err)
	   }
	*/

	return nil, nil
}

func (r *AccessControlPostgreSQL) ApiDelete(ctx context.Context, ctr *domain.ApiCriteria) error {

	/*
	   if customer.Id == "" {
	       return nil, errors.New("repository:mongodb: Update failed: customer id required")
	   }
	   qry := r.db.DB.WithContext(ctx)
	   if customer.Id != nil {
	       qry = qry.Where("id = ?", customer.Id)
	   }

	   if ctr.WithDeleted != nil {
	       if *ctr.WithDeleted {
	           qry = qry.Unscoped()
	       }
	   }

	   if err := qry.Delete(&domain.Budget{}).Error; err != nil {
	       return fmt.Errorf("repository:postgreSQL: failed to Delete budget: %v", err)
	   }

	   if qry.RowsAffected == 0 {
	       return errors.New("repository:postgreSQL: failed to Delete budget, 0 row affeted")
	   }
	*/

	return nil
}

func (r *AccessControlPostgreSQL) GroupStore(ctx context.Context, group *domain.Group) (*domain.Group, error) {
	q := r.db.DB

	err := q.WithContext(ctx).Create(group).Error
	if err != nil {
		return nil, fmt.Errorf("repository:postgreSQL: failed to Store group: %v", err)
	}

	return group, nil
}

func (r *AccessControlPostgreSQL) GroupGet(ctx context.Context, ctr *domain.GroupCriteria) (*domain.Group, error) {
	q := r.db.DB

	qry := q.WithContext(ctx)

	if ctr.ApiKey != "" {
		qry = qry.Where("api_key = ?", ctr.ApiKey)
	}

	var group domain.Group
	if err := qry.First(&group).Error; err != nil {
		return nil, fmt.Errorf("repository:postgreSQL: failed to FetchOne group: %v", err)
	}

	return &group, nil
}

func (r *AccessControlPostgreSQL) GroupList(ctx context.Context, ctr *domain.GroupCriteria) ([]*domain.Group, error) {

	qry := r.db.DB.WithContext(ctx) // operation on gorm db

	//if len(ctr.Ids) > 0 {
	//	qry = qry.Where("id IN(?)", ctr.Ids)
	//}
	//
	//if ctr.Id != nil {
	//	qry = qry.Where("id = ?", ctr.Id)
	//}
	//
	//if ctr.Phone != nil {
	//	qry = qry.Where("phone = ?", ctr.Phone)
	//}
	//
	//if ctr.TotalOrder != nil {
	//	qry = qry.Where("total_order = ?", ctr.TotalOrder)
	//}
	//
	//if ctr.SuccessfulOrder != nil {
	//	qry = qry.Where("successful_order = ?", ctr.SuccessfulOrder)
	//}
	//
	//if ctr.FraudOrder != nil {
	//	qry = qry.Where("fraud_order = ?", ctr.FraudOrder)
	//}
	//
	//if ctr.WithDeleted != nil {
	//	if *ctr.WithDeleted {
	//		qry = qry.Unscoped()
	//	}
	//}
	//
	//if ctr.Offset != nil {
	//	qry = qry.Offset(int(*ctr.Offset))
	//}
	//if ctr.Limit != nil {
	//	qry = qry.Limit(int(*ctr.Limit))
	//}
	//
	//if ctr.SortAsc {
	//	qry = qry.Order("id ASC")
	//} else {
	//	qry = qry.Order("id DESC")
	//}

	var apiList []*domain.Group

	if err := qry.Find(&apiList).Error; err != nil {
		return nil, fmt.Errorf("repository:postgreSQL: failed to Fetch customer: %v", err)
	}

	return apiList, nil
}

func (r *AccessControlPostgreSQL) GroupUpdate(ctx context.Context, group *domain.Group) (*domain.Group, error) {

	/*
	   if customer.Id == "" {
	       return nil, errors.New("repository:mongodb: Update failed: customer id required")
	   }

	   if err := r.db.DB.WithContext(ctx).Unscoped().Updates(customer).Error; err != nil {
	       return fmt.Errorf("repository:postgreSQLdb: Update failed: %v", err)
	   }
	*/

	return nil, nil
}

func (r *AccessControlPostgreSQL) GroupDelete(ctx context.Context, ctr *domain.GroupCriteria) error {

	/*
	   if customer.Id == "" {
	       return nil, errors.New("repository:mongodb: Update failed: customer id required")
	   }
	   qry := r.db.DB.WithContext(ctx)
	   if customer.Id != nil {
	       qry = qry.Where("id = ?", customer.Id)
	   }

	   if ctr.WithDeleted != nil {
	       if *ctr.WithDeleted {
	           qry = qry.Unscoped()
	       }
	   }

	   if err := qry.Delete(&domain.Budget{}).Error; err != nil {
	       return fmt.Errorf("repository:postgreSQL: failed to Delete budget: %v", err)
	   }

	   if qry.RowsAffected == 0 {
	       return errors.New("repository:postgreSQL: failed to Delete budget, 0 row affeted")
	   }
	*/

	return nil
}
