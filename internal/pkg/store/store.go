package store

type Store interface {
	Ping() error
	Grace()
}
