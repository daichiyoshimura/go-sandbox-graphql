package service

import (
	"context"
	"sandbox-gql/ent"
	"sandbox-gql/graph/model"
	"sandbox-gql/internal/mapping"
)

type AccountService struct {
	dbClient *ent.Client
}

func NewAccountService(dbClient *ent.Client) *AccountService {
	return &AccountService{
		dbClient: dbClient,
	}
}

func (s *AccountService) Create(ctx context.Context, input model.CreateAccountInput) (*model.Account, error) {
	entInput := mapping.ToEntCreateAccountInput(input)
	entAccount, err := s.dbClient.Account.Create().SetInput(entInput).Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapping.ToGraphAccount(entAccount), nil
}

func (s *AccountService) List(ctx context.Context) ([]*model.Account, error) {
	entAccounts, err := s.dbClient.Account.Query().WithItems().All(ctx)
	if err != nil {
		return nil, err
	}
	return mapping.ToGraphAccounts(entAccounts), nil
}
