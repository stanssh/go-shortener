package service

import (
	"crypto/rand"
	"strings"

	"github.com/stanssh/go-shortener/internal/storage"
)

type Service struct {
	Repository *storage.InMEM
}

var svc *Service

func SetStorage(s *storage.InMEM) {
	svc = &Service{Repository: s}
	// svc.Storage = s
}

func StoreURL(s string) (string, error) {
	myHash := GenerateString(s)
	svc.Repository.Put(
		s,
		myHash,
	)
	return myHash, nil
}

func RetrieveURL(s string) (string, error) {
	return svc.Repository.Get(s)
}

func GenerateString(s string) string {
	r := strings.ToLower(rand.Text()[:8])
	return r
}
