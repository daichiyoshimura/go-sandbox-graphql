package service

import (
	"context"
	"sandbox-gql/ent"
	"sandbox-gql/ent/account"
	"sandbox-gql/ent/customer"
	"sandbox-gql/graph/model"
	"sandbox-gql/internal/db"
	"sandbox-gql/internal/mapping"
)

type CustomerService struct {
	dbClient *ent.Client
}

func NewCustomerService(dbClient *ent.Client) *CustomerService {
	return &CustomerService{
		dbClient: dbClient,
	}
}

func (s *CustomerService) List(ctx context.Context) ([]*model.Customer, error) {
	entCustomers, err := s.dbClient.Customer.Query().WithFollows(func(q *ent.AccountQuery) {
		q.WithItems()
	}).All(ctx)
	if err != nil {
		return nil, err
	}
	return mapping.ToGraphCustomers(entCustomers), nil
}

func (s *CustomerService) Create(ctx context.Context, input model.CreateCustomerInput) (*model.Customer, error) {
	entInput := mapping.ToEntCreateCustomerInput(input)
	entCustomer, err := s.dbClient.Customer.Create().SetInput(entInput).Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapping.ToGraphCustomer(entCustomer), nil
}

func (s *CustomerService) Update(ctx context.Context, id int, input model.UpdateCustomerInput) (*model.Customer, error) {
	return db.RunInTransaction(ctx, s.dbClient, func(tx *ent.Tx) (*model.Customer, error) {

		customerUpdate := tx.Customer.UpdateOneID(id)
		if len(input.AddFollowIDs) > 0 {
			follows, err := tx.Account.Query().
				Where(account.IDIn(input.AddFollowIDs...)).
				All(ctx)
			if err != nil {
				return nil, err
			}
			customerUpdate.AddFollows(follows...)
		}

		if len(input.RemoveFollowIDs) > 0 {
			follows, err := tx.Account.Query().
				Where(account.IDIn(input.RemoveFollowIDs...)).
				All(ctx)
			if err != nil {
				return nil, err
			}
			customerUpdate.RemoveFollows(follows...)
		}

		entInput := mapping.ToEntUpdateCustomerInput(input)
		if _, err := customerUpdate.SetInput(entInput).Save(ctx); err != nil {
			return nil, err
		}

		entCustomer, err := tx.Customer.Query().
			Where(customer.ID(id)).
			WithFollows().
			Only(ctx)
		if err != nil {
			return nil, err
		}

		return mapping.ToGraphCustomer(entCustomer), nil
	})
}
