package service

import (
	"crypto/rand"
	"strings"

	"github.com/stanssh/go-shortener/internal/storage"
)

type Service struct {
	Storage *storage.InMEM
}

func New() *Service {
	return &Service{
		Storage: storage.NewMem(),
	}
}

func (svc *Service) StoreURL(s string) (string, error) {
	myHash := GenerateString(s)
	svc.Storage.Put(
		s,
		myHash,
	)
	return myHash, nil
}

func (svc *Service) RetrieveURL(s string) (string, error) {
	return svc.Storage.Get(s)
}

func GenerateString(s string) string {
	r := strings.ToLower(rand.Text()[:8])
	return r
}
