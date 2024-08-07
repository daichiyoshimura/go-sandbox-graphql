package mapping

import (
	"sandbox-gql/ent"
	"sandbox-gql/graph/model"
)

func ToGraphItem(item *ent.Item, account *model.Account) *model.Item {
	return &model.Item{
		ID:      item.ID,
		Name:    item.Name,
		Price:   item.Price,
		Account: account,
	}
}

func ToGraphItems(items []*ent.Item) []*model.Item {
	modelItems := make([]*model.Item, len(items))
	for i, entItem := range items {
		modelItems[i] = ToGraphItem(entItem, nil)
	}
	return modelItems
}

func ToEntCreateItemInput(input model.CreateItemInput) ent.CreateItemInput {
	return ent.CreateItemInput{
		Name:      input.Name,
		Price:     input.Price,
		AccountID: input.AccountID,
	}
}
