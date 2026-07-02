package image

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
CREATE TABLE IF NOT EXISTS images (
	sequence INTEGER PRIMARY KEY AUTOINCREMENT,
	id TEXT NOT NULL UNIQUE,
	name TEXT NOT NULL,
	status TEXT NOT NULL,
	container_format TEXT NOT NULL,
	disk_format TEXT NOT NULL,
	min_disk INTEGER NOT NULL,
	min_ram INTEGER NOT NULL,
	protected INTEGER NOT NULL,
	visibility TEXT NOT NULL,
	tags_json TEXT NOT NULL,
	created_at TEXT NOT NULL,
	updated_at TEXT NOT NULL
)`)

	return err
}

func (r *SQLiteRepository) Close() error {
	return r.db.Close()
}

func (r *SQLiteRepository) Create(image Image) Image {
	tagsJSON, err := json.Marshal(image.Tags)
	if err != nil {
		panic(fmt.Errorf("marshal image tags: %w", err))
	}

	_, err = r.db.Exec(
		`INSERT INTO images (
			id,
			name,
			status,
			container_format,
			disk_format,
			min_disk,
			min_ram,
			protected,
			visibility,
			tags_json,
			created_at,
			updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		image.ID,
		image.Name,
		image.Status,
		image.ContainerFormat,
		image.DiskFormat,
		image.MinDisk,
		image.MinRAM,
		boolToInt(image.Protected),
		image.Visibility,
		string(tagsJSON),
		image.CreatedAt,
		image.UpdatedAt,
	)
	if err != nil {
		panic(fmt.Errorf("insert image: %w", err))
	}

	return image
}

func (r *SQLiteRepository) List() []Image {
	rows, err := r.db.Query(`
SELECT
	id,
	name,
	status,
	container_format,
	disk_format,
	min_disk,
	min_ram,
	protected,
	visibility,
	tags_json,
	created_at,
	updated_at
FROM images
ORDER BY sequence`)
	if err != nil {
		panic(fmt.Errorf("list images: %w", err))
	}
	defer rows.Close()

	images := []Image{}
	for rows.Next() {
		image, err := scanImage(rows)
		if err != nil {
			panic(fmt.Errorf("scan image: %w", err))
		}
		images = append(images, image)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Errorf("iterate images: %w", err))
	}

	return images
}

func (r *SQLiteRepository) Get(id string) (Image, error) {
	row := r.db.QueryRow(`
SELECT
	id,
	name,
	status,
	container_format,
	disk_format,
	min_disk,
	min_ram,
	protected,
	visibility,
	tags_json,
	created_at,
	updated_at
FROM images
WHERE id = ?`, id)

	image, err := scanImage(row)
	if errors.Is(err, sql.ErrNoRows) {
		return Image{}, ErrImageNotFound
	}
	if err != nil {
		return Image{}, err
	}

	return image, nil
}

func (r *SQLiteRepository) Update(image Image) (Image, error) {
	tagsJSON, err := json.Marshal(image.Tags)
	if err != nil {
		return Image{}, fmt.Errorf("marshal image tags: %w", err)
	}

	result, err := r.db.Exec(
		`UPDATE images
		SET name = ?,
			status = ?,
			container_format = ?,
			disk_format = ?,
			min_disk = ?,
			min_ram = ?,
			protected = ?,
			visibility = ?,
			tags_json = ?,
			created_at = ?,
			updated_at = ?
		WHERE id = ?`,
		image.Name,
		image.Status,
		image.ContainerFormat,
		image.DiskFormat,
		image.MinDisk,
		image.MinRAM,
		boolToInt(image.Protected),
		image.Visibility,
		string(tagsJSON),
		image.CreatedAt,
		image.UpdatedAt,
		image.ID,
	)
	if err != nil {
		return Image{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return Image{}, err
	}
	if rowsAffected == 0 {
		return Image{}, ErrImageNotFound
	}

	return image, nil
}

func (r *SQLiteRepository) Delete(id string) error {
	result, err := r.db.Exec(`DELETE FROM images WHERE id = ?`, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrImageNotFound
	}

	return nil
}

func (r *SQLiteRepository) Reset() {
	if _, err := r.db.Exec(`DELETE FROM images`); err != nil {
		panic(fmt.Errorf("reset images: %w", err))
	}
}

type imageScanner interface {
	Scan(dest ...any) error
}

func scanImage(scanner imageScanner) (Image, error) {
	var image Image
	var protected int
	var tagsJSON string

	if err := scanner.Scan(
		&image.ID,
		&image.Name,
		&image.Status,
		&image.ContainerFormat,
		&image.DiskFormat,
		&image.MinDisk,
		&image.MinRAM,
		&protected,
		&image.Visibility,
		&tagsJSON,
		&image.CreatedAt,
		&image.UpdatedAt,
	); err != nil {
		return Image{}, err
	}

	if err := json.Unmarshal([]byte(tagsJSON), &image.Tags); err != nil {
		return Image{}, err
	}
	image.Protected = protected != 0

	return image, nil
}

func boolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}
