package identity

import (
	"database/sql"
	"errors"
	"fmt"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	_ "modernc.org/sqlite"
)

type SQLiteCatalogRepositories struct {
	db        *sql.DB
	Services  *SQLiteServiceRepository
	Endpoints *SQLiteEndpointRepository
}

func OpenSQLiteCatalogRepositories(
	dataSourceName string,
) (*SQLiteCatalogRepositories, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	repositories := &SQLiteCatalogRepositories{
		db:        db,
		Services:  &SQLiteServiceRepository{db: db},
		Endpoints: &SQLiteEndpointRepository{db: db},
	}
	if err := repositories.Init(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return repositories, nil
}

func (r *SQLiteCatalogRepositories) Init() error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS identity_services (
			sequence INTEGER PRIMARY KEY AUTOINCREMENT,
			id TEXT NOT NULL UNIQUE,
			name TEXT NOT NULL,
			type TEXT NOT NULL,
			description TEXT NOT NULL,
			enabled INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS identity_endpoints (
			sequence INTEGER PRIMARY KEY AUTOINCREMENT,
			id TEXT NOT NULL UNIQUE,
			service_id TEXT NOT NULL,
			interface TEXT NOT NULL,
			region TEXT NOT NULL,
			path TEXT NOT NULL,
			enabled INTEGER NOT NULL,
			description TEXT NOT NULL
		)`,
	}

	for _, statement := range statements {
		if _, err := r.db.Exec(statement); err != nil {
			return err
		}
	}

	return nil
}

func (r *SQLiteCatalogRepositories) Close() error {
	return r.db.Close()
}

type SQLiteServiceRepository struct {
	db *sql.DB
}

func (r *SQLiteServiceRepository) Save(
	service appidentity.ServiceDefinition,
) appidentity.ServiceDefinition {
	_, err := r.db.Exec(
		`INSERT INTO identity_services (
			id,
			name,
			type,
			description,
			enabled
		) VALUES (?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			name = excluded.name,
			type = excluded.type,
			description = excluded.description,
			enabled = excluded.enabled`,
		service.ID,
		service.Name,
		service.Type,
		service.Description,
		boolToInt(service.Enabled),
	)
	if err != nil {
		panic(fmt.Errorf("save identity service: %w", err))
	}

	return service
}

func (r *SQLiteServiceRepository) List() []appidentity.ServiceDefinition {
	rows, err := r.db.Query(`
SELECT
	id,
	name,
	type,
	description,
	enabled
FROM identity_services
ORDER BY sequence`)
	if err != nil {
		panic(fmt.Errorf("list identity services: %w", err))
	}
	defer rows.Close()

	services := []appidentity.ServiceDefinition{}
	for rows.Next() {
		service, err := scanService(rows)
		if err != nil {
			panic(fmt.Errorf("scan identity service: %w", err))
		}
		services = append(services, service)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Errorf("iterate identity services: %w", err))
	}

	return services
}

func (r *SQLiteServiceRepository) Get(
	id string,
) (appidentity.ServiceDefinition, error) {
	row := r.db.QueryRow(`
SELECT
	id,
	name,
	type,
	description,
	enabled
FROM identity_services
WHERE id = ?`, id)

	service, err := scanService(row)
	if errors.Is(err, sql.ErrNoRows) {
		return appidentity.ServiceDefinition{}, appidentity.ErrServiceNotFound
	}
	if err != nil {
		return appidentity.ServiceDefinition{}, err
	}

	return service, nil
}

func (r *SQLiteServiceRepository) Reset() {
	if _, err := r.db.Exec(`DELETE FROM identity_services`); err != nil {
		panic(fmt.Errorf("reset identity services: %w", err))
	}
}

type SQLiteEndpointRepository struct {
	db *sql.DB
}

func (r *SQLiteEndpointRepository) Save(
	endpoint appidentity.EndpointDefinition,
) appidentity.EndpointDefinition {
	_, err := r.db.Exec(
		`INSERT INTO identity_endpoints (
			id,
			service_id,
			interface,
			region,
			path,
			enabled,
			description
		) VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			service_id = excluded.service_id,
			interface = excluded.interface,
			region = excluded.region,
			path = excluded.path,
			enabled = excluded.enabled,
			description = excluded.description`,
		endpoint.ID,
		endpoint.ServiceID,
		endpoint.Interface,
		endpoint.Region,
		endpoint.Path,
		boolToInt(endpoint.Enabled),
		endpoint.Description,
	)
	if err != nil {
		panic(fmt.Errorf("save identity endpoint: %w", err))
	}

	return endpoint
}

func (r *SQLiteEndpointRepository) List() []appidentity.EndpointDefinition {
	rows, err := r.db.Query(`
SELECT
	id,
	service_id,
	interface,
	region,
	path,
	enabled,
	description
FROM identity_endpoints
ORDER BY sequence`)
	if err != nil {
		panic(fmt.Errorf("list identity endpoints: %w", err))
	}
	defer rows.Close()

	endpoints := []appidentity.EndpointDefinition{}
	for rows.Next() {
		endpoint, err := scanEndpoint(rows)
		if err != nil {
			panic(fmt.Errorf("scan identity endpoint: %w", err))
		}
		endpoints = append(endpoints, endpoint)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Errorf("iterate identity endpoints: %w", err))
	}

	return endpoints
}

func (r *SQLiteEndpointRepository) Get(
	id string,
) (appidentity.EndpointDefinition, error) {
	row := r.db.QueryRow(`
SELECT
	id,
	service_id,
	interface,
	region,
	path,
	enabled,
	description
FROM identity_endpoints
WHERE id = ?`, id)

	endpoint, err := scanEndpoint(row)
	if errors.Is(err, sql.ErrNoRows) {
		return appidentity.EndpointDefinition{}, appidentity.ErrEndpointNotFound
	}
	if err != nil {
		return appidentity.EndpointDefinition{}, err
	}

	return endpoint, nil
}

func (r *SQLiteEndpointRepository) ListByServiceID(
	serviceID string,
) []appidentity.EndpointDefinition {
	rows, err := r.db.Query(`
SELECT
	id,
	service_id,
	interface,
	region,
	path,
	enabled,
	description
FROM identity_endpoints
WHERE service_id = ?
ORDER BY sequence`, serviceID)
	if err != nil {
		panic(fmt.Errorf("list identity endpoints by service: %w", err))
	}
	defer rows.Close()

	endpoints := []appidentity.EndpointDefinition{}
	for rows.Next() {
		endpoint, err := scanEndpoint(rows)
		if err != nil {
			panic(fmt.Errorf("scan identity endpoint: %w", err))
		}
		endpoints = append(endpoints, endpoint)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Errorf("iterate identity endpoints: %w", err))
	}

	return endpoints
}

func (r *SQLiteEndpointRepository) Reset() {
	if _, err := r.db.Exec(`DELETE FROM identity_endpoints`); err != nil {
		panic(fmt.Errorf("reset identity endpoints: %w", err))
	}
}

func scanService(scanner rowScanner) (appidentity.ServiceDefinition, error) {
	var service appidentity.ServiceDefinition
	var enabled int
	err := scanner.Scan(
		&service.ID,
		&service.Name,
		&service.Type,
		&service.Description,
		&enabled,
	)
	if err != nil {
		return appidentity.ServiceDefinition{}, err
	}
	service.Enabled = intToBool(enabled)

	return service, nil
}

func scanEndpoint(scanner rowScanner) (appidentity.EndpointDefinition, error) {
	var endpoint appidentity.EndpointDefinition
	var enabled int
	err := scanner.Scan(
		&endpoint.ID,
		&endpoint.ServiceID,
		&endpoint.Interface,
		&endpoint.Region,
		&endpoint.Path,
		&enabled,
		&endpoint.Description,
	)
	if err != nil {
		return appidentity.EndpointDefinition{}, err
	}
	endpoint.Enabled = intToBool(enabled)

	return endpoint, nil
}
