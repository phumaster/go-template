package configs

// DB struct
type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

// GetConfig return the configs
func (d *DB) GetConfig() *DB {
	return &DB{
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "root",
		Password: "example",
		DBName:   "go-template",
	}
}
