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

```zsh
go run -mod=mod entgo.io/ent/cmd/ent new Todo
```