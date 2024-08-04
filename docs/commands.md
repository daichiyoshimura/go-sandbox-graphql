# Commands for building this repository

## install graphql

```zsh
go install github.com/99designs/gqlgen@latest
```

- check version

```zsh
gqlgen version
```

## init go project

```zsh
go mod init sandbox-gql
go get -u github.com/99designs/gqlgen
gqlgen init
```
