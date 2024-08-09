package service_test

import (
	"context"
	"encoding/json"
	"sandbox-gql/ent/enttest"
	"sandbox-gql/graph/model"
	"sandbox-gql/internal/service"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"github.com/google/go-cmp/cmp"
)

func TestAccountService_Create(t *testing.T) {

	tests := []struct {
		name    string
		input   model.CreateAccountInput
		want    *model.Account
		wantErr error
	}{
		{
			"add positive numbers",
			model.CreateAccountInput{
				Name:  "John",
				Email: "example.com",
			},
			&model.Account{
				ID:    1,
				Name:  "John",
				Email: "example.com",
			},
			nil,
		},
	}

	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	service := service.NewAccountService(client)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := service.Create(context.Background(), tt.input)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("got: %s; want: %s", err.Error(), tt.wantErr.Error())
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				gotStr, _ := json.Marshal(got)
				wantStr, _ := json.Marshal(tt.want)
				t.Errorf("Mismatch got:%s want:%s :%s", gotStr, wantStr, diff)
			}
		})
	}

	client.Account.Delete().ExecX(context.Background())
}
