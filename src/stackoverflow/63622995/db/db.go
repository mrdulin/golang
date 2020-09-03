package db

type OrmDBWithError struct {
	OrmDB
	Error error
}

type OrmDB interface {
	Table(name string) OrmDB
	Where(query interface{}, args ...interface{}) OrmDB
	Scan(dest interface{}) *OrmDBWithError
}
