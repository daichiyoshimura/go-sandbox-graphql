package db

import (
	"sandbox-gql/ent"

	"entgo.io/ent/dialect"
)

func Client() (*ent.Client, error) {
	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())
	client, err := ent.Open(dialect.MySQL, dsn("usr", "usrpwd", "db", "3306", "db"), entOptions...)
	if err != nil {
		return nil, err
	}
	defer client.Close()
	return client, nil
}
