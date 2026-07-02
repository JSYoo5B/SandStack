package identity

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	_ "modernc.org/sqlite"
)

type SQLiteTokenRepository struct {
	db *sql.DB
}

func OpenSQLiteTokenRepository(
	dataSourceName string,
) (*SQLiteTokenRepository, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	repository := &SQLiteTokenRepository{db: db}
	if err := repository.Init(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return repository, nil
}

func (r *SQLiteTokenRepository) Init() error {
	_, err := r.db.Exec(`
CREATE TABLE IF NOT EXISTS identity_tokens (
	sequence INTEGER PRIMARY KEY AUTOINCREMENT,
	id TEXT NOT NULL UNIQUE,
	token_json TEXT NOT NULL
)`)

	return err
}

func (r *SQLiteTokenRepository) Close() error {
	return r.db.Close()
}

func (r *SQLiteTokenRepository) Save(
	token appidentity.IssuedToken,
) appidentity.IssuedToken {
	tokenJSON, err := json.Marshal(token)
	if err != nil {
		panic(fmt.Errorf("marshal identity token: %w", err))
	}

	_, err = r.db.Exec(
		`INSERT INTO identity_tokens (
			id,
			token_json
		) VALUES (?, ?)
		ON CONFLICT(id) DO UPDATE SET
			token_json = excluded.token_json`,
		token.ID,
		string(tokenJSON),
	)
	if err != nil {
		panic(fmt.Errorf("save identity token: %w", err))
	}

	return token
}

func (r *SQLiteTokenRepository) Get(
	id string,
) (appidentity.IssuedToken, error) {
	row := r.db.QueryRow(`
SELECT token_json
FROM identity_tokens
WHERE id = ?`, id)

	token, err := scanToken(row)
	if errors.Is(err, sql.ErrNoRows) {
		return appidentity.IssuedToken{}, appidentity.ErrTokenNotFound
	}
	if err != nil {
		return appidentity.IssuedToken{}, err
	}

	return token, nil
}

func (r *SQLiteTokenRepository) Delete(id string) error {
	result, err := r.db.Exec(`DELETE FROM identity_tokens WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return appidentity.ErrTokenNotFound
	}

	return nil
}

func (r *SQLiteTokenRepository) Reset() {
	if _, err := r.db.Exec(`DELETE FROM identity_tokens`); err != nil {
		panic(fmt.Errorf("reset identity tokens: %w", err))
	}
}

func scanToken(scanner rowScanner) (appidentity.IssuedToken, error) {
	var tokenJSON string
	if err := scanner.Scan(&tokenJSON); err != nil {
		return appidentity.IssuedToken{}, err
	}

	var token appidentity.IssuedToken
	if err := json.Unmarshal([]byte(tokenJSON), &token); err != nil {
		return appidentity.IssuedToken{}, err
	}

	return token, nil
}
