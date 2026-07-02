package compute

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

type SQLiteServerRepository struct {
	db *sql.DB
}

func OpenSQLiteServerRepository(
	dataSourceName string,
) (*SQLiteServerRepository, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	repository := &SQLiteServerRepository{db: db}
	if err := repository.Init(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return repository, nil
}

func (r *SQLiteServerRepository) Init() error {
	_, err := r.db.Exec(`
CREATE TABLE IF NOT EXISTS compute_servers (
	sequence INTEGER PRIMARY KEY AUTOINCREMENT,
	id TEXT NOT NULL UNIQUE,
	name TEXT NOT NULL,
	image_id TEXT NOT NULL,
	flavor_id TEXT NOT NULL,
	tenant_id TEXT NOT NULL,
	user_id TEXT NOT NULL,
	status TEXT NOT NULL,
	progress INTEGER NOT NULL,
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL,
	metadata_json TEXT NOT NULL
)`)

	return err
}

func (r *SQLiteServerRepository) Close() error {
	return r.db.Close()
}

func (r *SQLiteServerRepository) Create(server Server) Server {
	metadataJSON := marshalStringMap(server.Metadata)

	_, err := r.db.Exec(
		`INSERT INTO compute_servers (
			id,
			name,
			image_id,
			flavor_id,
			tenant_id,
			user_id,
			status,
			progress,
			created_at,
			updated_at,
			metadata_json
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		server.ID,
		server.Name,
		server.ImageID,
		server.FlavorID,
		server.TenantID,
		server.UserID,
		server.Status,
		server.Progress,
		server.CreatedAt,
		server.UpdatedAt,
		metadataJSON,
	)
	if err != nil {
		panic(fmt.Errorf("insert server: %w", err))
	}

	return server
}

func (r *SQLiteServerRepository) List() []Server {
	rows, err := r.db.Query(`
SELECT
	id,
	name,
	image_id,
	flavor_id,
	tenant_id,
	user_id,
	status,
	progress,
	created_at,
	updated_at,
	metadata_json
FROM compute_servers
ORDER BY sequence`)
	if err != nil {
		panic(fmt.Errorf("list servers: %w", err))
	}
	defer rows.Close()

	servers := []Server{}
	for rows.Next() {
		server, err := scanServer(rows)
		if err != nil {
			panic(fmt.Errorf("scan server: %w", err))
		}
		servers = append(servers, server)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Errorf("iterate servers: %w", err))
	}

	return servers
}

func (r *SQLiteServerRepository) Get(id string) (Server, error) {
	row := r.db.QueryRow(`
SELECT
	id,
	name,
	image_id,
	flavor_id,
	tenant_id,
	user_id,
	status,
	progress,
	created_at,
	updated_at,
	metadata_json
FROM compute_servers
WHERE id = ?`, id)

	server, err := scanServer(row)
	if errors.Is(err, sql.ErrNoRows) {
		return Server{}, ErrServerNotFound
	}
	if err != nil {
		return Server{}, err
	}

	return server, nil
}

func (r *SQLiteServerRepository) Update(server Server) (Server, error) {
	metadataJSON := marshalStringMap(server.Metadata)
	result, err := r.db.Exec(
		`UPDATE compute_servers
		SET name = ?,
			image_id = ?,
			flavor_id = ?,
			tenant_id = ?,
			user_id = ?,
			status = ?,
			progress = ?,
			created_at = ?,
			updated_at = ?,
			metadata_json = ?
		WHERE id = ?`,
		server.Name,
		server.ImageID,
		server.FlavorID,
		server.TenantID,
		server.UserID,
		server.Status,
		server.Progress,
		server.CreatedAt,
		server.UpdatedAt,
		metadataJSON,
		server.ID,
	)
	if err != nil {
		return Server{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Server{}, err
	}
	if rowsAffected == 0 {
		return Server{}, ErrServerNotFound
	}

	return server, nil
}

func (r *SQLiteServerRepository) Delete(id string) error {
	result, err := r.db.Exec(`DELETE FROM compute_servers WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrServerNotFound
	}

	return nil
}

func (r *SQLiteServerRepository) Reset() {
	if _, err := r.db.Exec(`DELETE FROM compute_servers`); err != nil {
		panic(fmt.Errorf("reset servers: %w", err))
	}
}

type serverScanner interface {
	Scan(dest ...any) error
}

func scanServer(scanner serverScanner) (Server, error) {
	var server Server
	var metadataJSON string

	if err := scanner.Scan(
		&server.ID,
		&server.Name,
		&server.ImageID,
		&server.FlavorID,
		&server.TenantID,
		&server.UserID,
		&server.Status,
		&server.Progress,
		&server.CreatedAt,
		&server.UpdatedAt,
		&metadataJSON,
	); err != nil {
		return Server{}, err
	}

	if err := json.Unmarshal([]byte(metadataJSON), &server.Metadata); err != nil {
		return Server{}, err
	}

	return server, nil
}

func marshalStringMap(value map[string]string) string {
	data, err := json.Marshal(value)
	if err != nil {
		panic(fmt.Errorf("marshal string map: %w", err))
	}

	return string(data)
}
