package service

type Repository interface {
	Put(string, string) (string, error)
	Get(string) (string, error)
}
