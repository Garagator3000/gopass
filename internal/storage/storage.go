package storage

import (
	"context"

	"github.com/garagator3000/gopass/internal/entities"
)

type Storage interface {
	CreateSecret(ctx context.Context, secret entities.Secret) error
	ReadSecret(ctx context.Context, name string) (string, error)
	ListSecret(ctx context.Context, groupname string) ([]string, error)

	Close()
}

func Init(storageType, storagePath string) Storage {
	switch storageType {
	case "sqlite":
		return NewSqlite(storagePath)
	default:
		return NewSqlite(storagePath)
	}
}
