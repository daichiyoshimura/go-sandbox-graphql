package mapping

import (
	"sandbox-gql/ent"
	"sandbox-gql/graph/model"
)

func ToGraphAccount(account *ent.Account) *model.Account {
	if account == nil {
		return nil
	}
	return &model.Account{
		ID:    account.ID,
		Name:  account.Name,
		Email: account.Email,
		Items: ToGraphItems(account.Edges.Items),
	}
}

func ToGraphAccounts(accounts []*ent.Account) []*model.Account {
	if accounts == nil {
		return nil
	}
	modelAccounts := make([]*model.Account, len(accounts))
	for i, entAccount := range accounts {
		modelAccounts[i] = ToGraphAccount(entAccount)
	}
	return modelAccounts
}

func ToEntCreateAccountInput(input model.CreateAccountInput) ent.CreateAccountInput {
	return ent.CreateAccountInput{
		Name:  input.Name,
		Email: input.Email,
	}
}
