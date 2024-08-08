package service

import "sandbox-gql/ent"

type iDBClient interface {
	ent.Rollbacker
	ent.Committer
	ent.Querier
	ent.Mutator
}
