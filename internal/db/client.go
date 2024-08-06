package db

import (
	"fmt"
	"sandbox-gql/ent"
	"sandbox-gql/internal/env"

	"entgo.io/ent/dialect"
)

func Client(dbvars *env.DB) (*ent.Client, error) {
	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())
	return ent.Open(dialect.MySQL, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbvars.User(), dbvars.Password(), dbvars.Host(), dbvars.Port(), dbvars.Name()), entOptions...)
}
