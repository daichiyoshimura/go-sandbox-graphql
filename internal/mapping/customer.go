package mapping

import (
	"sandbox-gql/ent"
	"sandbox-gql/graph/model"
)

func ToGraphCustomer(Customer *ent.Customer) *model.Customer {
	if Customer == nil {
		return nil
	}
	return &model.Customer{
		ID:      Customer.ID,
		Name:    Customer.Name,
		Email:   Customer.Email,
		Follows: ToGraphAccounts(Customer.Edges.Follows),
	}
}

func ToGraphCustomers(Customers []*ent.Customer) []*model.Customer {
	if Customers == nil {
		return nil
	}
	modelCustomers := make([]*model.Customer, len(Customers))
	for i, entCustomer := range Customers {
		modelCustomers[i] = ToGraphCustomer(entCustomer)
	}
	return modelCustomers
}

func ToEntCreateCustomerInput(input model.CreateCustomerInput) ent.CreateCustomerInput {
	return ent.CreateCustomerInput{
		Name:  input.Name,
		Email: input.Email,
	}
}

func ToEntUpdateCustomerInput(input model.UpdateCustomerInput) ent.UpdateCustomerInput {
	return ent.UpdateCustomerInput{
		Name:  input.Name,
		Email: input.Email,
	}
}
