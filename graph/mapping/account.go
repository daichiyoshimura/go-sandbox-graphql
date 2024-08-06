package mapping

import (
	"sandbox-gql/ent"
	"sandbox-gql/graph/model"
)

func ToGraphAccount(account *ent.Account, items []*model.Item) *model.Account {
	return &model.Account{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
		Items: items,
	}
}

func ToEntCreateAccountInput(input model.CreateAccountInput) ent.CreateAccountInput {
	return ent.CreateAccountInput{
		Name:  input.Name,
		Email: input.Email,
	}
}
