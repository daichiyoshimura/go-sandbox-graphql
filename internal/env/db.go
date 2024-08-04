package env

type DB struct {
	host     string
	port     string
	user     string
	password string
	name     string
}

func loadDB() (*DB, error) {
	host, err := load("DB_HOST")
	if err != nil {
		return nil, err
	}

	port, err := load("DB_PORT")
	if err != nil {
		return nil, err
	}

	user, err := load("DB_USER")
	if err != nil {
		return nil, err
	}

	password, err := load("DB_PASSWORD")
	if err != nil {
		return nil, err
	}

	name, err := load("DB_NAME")
	if err != nil {
		return nil, err
	}

	return &DB{
		host:     host,
		port:     port,
		user:     user,
		password: password,
		name:     name,
	}, nil
}

func (db *DB) Host() string {
	return db.host
}

func (db *DB) Port() string {
	return db.port
}

func (db *DB) User() string {
	return db.user
}

func (db *DB) Password() string {
	return db.password
}

func (db *DB) Name() string {
	return db.name
}
