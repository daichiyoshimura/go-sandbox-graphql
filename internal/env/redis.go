package env

type Redis struct {
	host     string
	port     string
	password string
	name     string
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

	name, err := load("REDIS_NAME")
	if err != nil {
		return nil, err
	}

	return &Redis{
		host:     host,
		port:     port,
		password: password,
		name:     name,
	}, nil
}

func (db *Redis) Host() string {
	return db.host
}

func (db *Redis) Port() string {
	return db.port
}

func (db *Redis) Password() string {
	return db.password
}

func (db *Redis) Name() string {
	return db.name
}
