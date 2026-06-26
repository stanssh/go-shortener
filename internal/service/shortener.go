package service

import (
	"crypto/rand"
	"strings"

	"github.com/stanssh/go-shortener/internal/repository"
)

type Service struct {
	Storage repository.InMEM
}

func (svc *Service) StoreURL(s string) (string, error) {
	svc.Storage.Put(
		s,
		GenerateString(s),
	)
	return "", nil
}

func (svc *Service) RetrieveURL(s string) (string, error) {
	return svc.Storage.Get(s)
}

func GenerateString(s string) string {
	r := strings.ToLower(rand.Text()[:8])
	return r
}
