package volume

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	_ "modernc.org/sqlite"
)

type SQLiteRepository struct {
	db *sql.DB
}

func OpenSQLiteRepository(dataSourceName string) (*SQLiteRepository, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	repository := &SQLiteRepository{db: db}
	if err := repository.Init(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return repository, nil
}

func (r *SQLiteRepository) Init() error {
	_, err := r.db.Exec(`
CREATE TABLE IF NOT EXISTS volumes (
	sequence INTEGER PRIMARY KEY AUTOINCREMENT,
	id TEXT NOT NULL UNIQUE,
	status TEXT NOT NULL,
	size INTEGER NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	volume_type TEXT NOT NULL,
	metadata_json TEXT NOT NULL,
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL,
	bootable TEXT NOT NULL,
	encrypted INTEGER NOT NULL,
	multiattach INTEGER NOT NULL
)`)

	return err
}

func (r *SQLiteRepository) Close() error {
	return r.db.Close()
}

func (r *SQLiteRepository) Create(volume Volume) Volume {
	metadataJSON := marshalStringMap(volume.Metadata)
	_, err := r.db.Exec(
		`INSERT INTO volumes (
			id,
			status,
			size,
			name,
			description,
			volume_type,
			metadata_json,
			created_at,
			updated_at,
			bootable,
			encrypted,
			multiattach
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		volume.ID,
		volume.Status,
		volume.Size,
		volume.Name,
		volume.Description,
		volume.VolumeType,
		metadataJSON,
		volume.CreatedAt,
		volume.UpdatedAt,
		volume.Bootable,
		boolToInt(volume.Encrypted),
		boolToInt(volume.Multiattach),
	)
	if err != nil {
		panic(fmt.Errorf("insert volume: %w", err))
	}

	return volume
}

func (r *SQLiteRepository) List() []Volume {
	rows, err := r.db.Query(`
SELECT
	id,
	status,
	size,
	name,
	description,
	volume_type,
	metadata_json,
	created_at,
	updated_at,
	bootable,
	encrypted,
	multiattach
FROM volumes
ORDER BY sequence`)
	if err != nil {
		panic(fmt.Errorf("list volumes: %w", err))
	}
	defer rows.Close()

	volumes := []Volume{}
	for rows.Next() {
		volume, err := scanVolume(rows)
		if err != nil {
			panic(fmt.Errorf("scan volume: %w", err))
		}
		volumes = append(volumes, volume)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Errorf("iterate volumes: %w", err))
	}

	return volumes
}

func (r *SQLiteRepository) Get(id string) (Volume, error) {
	row := r.db.QueryRow(`
SELECT
	id,
	status,
	size,
	name,
	description,
	volume_type,
	metadata_json,
	created_at,
	updated_at,
	bootable,
	encrypted,
	multiattach
FROM volumes
WHERE id = ?`, id)

	volume, err := scanVolume(row)
	if errors.Is(err, sql.ErrNoRows) {
		return Volume{}, ErrVolumeNotFound
	}
	if err != nil {
		return Volume{}, err
	}

	return volume, nil
}

func (r *SQLiteRepository) Update(volume Volume) (Volume, error) {
	metadataJSON := marshalStringMap(volume.Metadata)
	result, err := r.db.Exec(
		`UPDATE volumes
		SET status = ?,
			size = ?,
			name = ?,
			description = ?,
			volume_type = ?,
			metadata_json = ?,
			created_at = ?,
			updated_at = ?,
			bootable = ?,
			encrypted = ?,
			multiattach = ?
		WHERE id = ?`,
		volume.Status,
		volume.Size,
		volume.Name,
		volume.Description,
		volume.VolumeType,
		metadataJSON,
		volume.CreatedAt,
		volume.UpdatedAt,
		volume.Bootable,
		boolToInt(volume.Encrypted),
		boolToInt(volume.Multiattach),
		volume.ID,
	)
	if err != nil {
		return Volume{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Volume{}, err
	}
	if rowsAffected == 0 {
		return Volume{}, ErrVolumeNotFound
	}

	return volume, nil
}

func (r *SQLiteRepository) Delete(id string) error {
	result, err := r.db.Exec(`DELETE FROM volumes WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrVolumeNotFound
	}

	return nil
}

func (r *SQLiteRepository) Reset() {
	if _, err := r.db.Exec(`DELETE FROM volumes`); err != nil {
		panic(fmt.Errorf("reset volumes: %w", err))
	}
}

type volumeScanner interface {
	Scan(dest ...any) error
}

func scanVolume(scanner volumeScanner) (Volume, error) {
	var volume Volume
	var metadataJSON string
	var encrypted int
	var multiattach int

	if err := scanner.Scan(
		&volume.ID,
		&volume.Status,
		&volume.Size,
		&volume.Name,
		&volume.Description,
		&volume.VolumeType,
		&metadataJSON,
		&volume.CreatedAt,
		&volume.UpdatedAt,
		&volume.Bootable,
		&encrypted,
		&multiattach,
	); err != nil {
		return Volume{}, err
	}

	if err := json.Unmarshal([]byte(metadataJSON), &volume.Metadata); err != nil {
		return Volume{}, err
	}
	volume.Encrypted = encrypted != 0
	volume.Multiattach = multiattach != 0

	return volume, nil
}

func marshalStringMap(value map[string]string) string {
	data, err := json.Marshal(value)
	if err != nil {
		panic(fmt.Errorf("marshal string map: %w", err))
	}

	return string(data)
}

func boolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}
