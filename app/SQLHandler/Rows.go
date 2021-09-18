package SQLHandler

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}
