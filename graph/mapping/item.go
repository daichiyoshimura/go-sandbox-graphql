package mapping

import (
	"sandbox-gql/ent"
	"sandbox-gql/graph/model"
)

func ToGraphItem(item *ent.Item) *model.Item {
	if item == nil {
		return nil
	}
	return &model.Item{
		ID:      item.ID,
		Name:    item.Name,
		Price:   item.Price,
		Account: ToGraphAccount(item.Edges.Account),
	}
}

func ToGraphItems(items []*ent.Item) []*model.Item {
	if items == nil {
		return nil
	}
	modelItems := make([]*model.Item, len(items))
	for i, entItem := range items {
		modelItems[i] = ToGraphItem(entItem)
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
