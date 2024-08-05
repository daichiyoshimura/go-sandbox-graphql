# go-graphql

## How to create new schema

### Step1. Generate ent code

```zsh
go run -mod=mod entgo.io/ent/cmd/ent new ${new-schema}
```

### Step2. Implement the schema

See `ent/schema/${new-schema}.go`

### Step3. Generate graph code

```zsh
go generate ./ent
```

### Step4. Implement resolvers

See `graph/resolver.go, graph/ent.resolver.go`