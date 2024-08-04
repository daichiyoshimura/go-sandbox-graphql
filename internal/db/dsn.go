package db

import "fmt"

func dsn(user, password, server, port, db string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, server, port, db)
}
