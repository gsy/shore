package database

type Database interface {
	Connect(dsn string) (db Database, err error)
}
