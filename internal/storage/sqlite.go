package storage

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/garagator3000/gopass/internal/entities"

	// Create extension module for sqlite for database/sql.
	_ "github.com/mattn/go-sqlite3"
)

const sqliteDriverName = "sqlite3"

type Sqlite struct {
	db *sql.DB
}

const createSecretTable = `
CREATE TABLE IF NOT EXISTS secret(
secret_id INTEGER PRIMARY KEY AUTOINCREMENT,
name TEXT,
data TEXT,
user TEXT,
created_at INTEGER, 
updated_at INTEGER);
`

func NewSqlite(dbPath string) *Sqlite {
	dbPath, err := setDBPath(dbPath)
	if err != nil {
		panic(fmt.Errorf("failed to set database path: %w", err))
	}

	db, err := sql.Open(sqliteDriverName, dbPath)
	if err != nil {
		panic(fmt.Errorf("failed to open sqlite db: %w", err))
	}

	if _, err = db.Exec(createSecretTable); err != nil {
		panic(fmt.Errorf("failed to create db table: %w", err))
	}

	return &Sqlite{
		db: db,
	}
}

//nolint:gosec // Just SQL query template.
const createSecretSQLite = `
INSERT INTO secret (name, data, user, created_at, updated_at)
VALUES (?, ?, ?, ?, ?);
`

func (s *Sqlite) CreateSecret(_ context.Context, secret entities.Secret) error {
	createdAt := secret.CreatedAt.UnixNano()
	updatedAt := secret.UpdatedAt.UnixNano()

	_, err := s.db.Exec(createSecretSQLite,
		secret.Name,
		secret.Data,
		secret.User,
		createdAt,
		updatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create secret %s: %w", secret.Name, err)
	}

	return nil
}

//nolint:gosec // Just SQL query template.
const readSecretSQLite = `
SELECT data
FROM secret
WHERE name = ?;
`

func (s *Sqlite) ReadSecret(_ context.Context, name string) (string, error) {
	var data string

	row := s.db.QueryRow(readSecretSQLite, name)

	err := row.Scan(&data)
	if err != nil {
		return "", fmt.Errorf("failed to read secret %s: %w", name, err)
	}

	return data, nil
}

func (s *Sqlite) Close() {
	s.db.Close()
}

func setDBPath(dbPath string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to determinate home dir: %w", err)
	}
	dirpath := home + "/.gopass"

	if dbPath != "" {
		return dbPath, nil
	}

	if mkdirErr := os.MkdirAll(dirpath, os.ModeDir|os.ModeAppend|os.ModePerm); mkdirErr != nil {
		return "", fmt.Errorf("faied to create directory %s: %w", dirpath, mkdirErr)
	}

	dbPath = dirpath + "/gopass.sqlite"

	return dbPath, nil
}
