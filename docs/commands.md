# Commands for building this repository

## install graphql

```zsh
go install github.com/99designs/gqlgen@latest
```

- check version

```zsh
gqlgen version
```

## init project

```zsh
go mod init sandbox-gql
go get -u github.com/99designs/gqlgen
gqlgen init
```

## generate schema

- 0. edit schema.graphql 
- 0. run this command

```zsh
gqlgen generate
```

## install ent

```zsh
go get entgo.io/ent/cmd/ent
go get entgo.io/contrib/entgql
```

## create shcelton code



## flow of creating new schema

- 0. generate schelton code for ent

```zsh
go run -mod=mod entgo.io/ent/cmd/ent new ${new-schema}
```

- 0. implement `ent/schema/${new-schema}`

```go
package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	// "entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email").
			NotEmpty().
			MaxLen(255),
		field.Bool("done").
			Default(false),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}

func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
```

- 0. editing ent/entc.go

```go
entgql.WithSchemaPath("../graph/todo.graphqls"),
```

- 0. run go generate 

```zsh
go generate ./ent
```