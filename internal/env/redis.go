package env

import "strconv"

type Redis struct {
	host     string
	port     string
	password string
	db       int
}

func loadRedis() (*Redis, error) {
	host, err := load("REDIS_HOST")
	if err != nil {
		return nil, err
	}

	port, err := load("REDIS_PORT")
	if err != nil {
		return nil, err
	}

	password, err := load("REDIS_PASSWORD")
	if err != nil {
		return nil, err
	}

	db, err := load("REDIS_DB")
	if err != nil {
		return nil, err
	}
	intDb, err := strconv.Atoi(db)
	if err != nil {
		return nil, err
	}

	return &Redis{
		host:     host,
		port:     port,
		password: password,
		db:       intDb,
	}, nil
}

func (r *Redis) Host() string {
	return r.host
}

func (r *Redis) Port() string {
	return r.port
}

func (r *Redis) Password() string {
	return r.password
}

func (r *Redis) DB() int {
	return r.db
}
