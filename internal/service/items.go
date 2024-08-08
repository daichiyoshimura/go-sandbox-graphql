package service

import (
	"context"
	"sandbox-gql/ent"
	"sandbox-gql/graph/model"
	"sandbox-gql/internal/mapping"
)

type ItemService struct {
	dbClient *ent.Client
}

func NewItemsService(dbClient *ent.Client) *ItemService {
	return &ItemService{
		dbClient: dbClient,
	}
}

func (s *ItemService) Create(ctx context.Context, input model.CreateItemInput) (*model.Item, error) {
	entInput := mapping.ToEntCreateItemInput(input)
	entItem, err := s.dbClient.Item.Create().SetInput(entInput).Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapping.ToGraphItem(entItem), nil
}

func (s *ItemService) List(ctx context.Context) ([]*model.Item, error) {
	entItems, err := s.dbClient.Item.Query().WithOwner().All(ctx)
	if err != nil {
		return nil, err
	}
	return mapping.ToGraphItems(entItems), nil
}
