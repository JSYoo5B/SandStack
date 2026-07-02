package identity

import (
	"database/sql"
	"errors"
	"fmt"

	appidentity "github.com/JSYoo5B/SandStack/internal/app/identity"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/projects"
	"github.com/gophercloud/gophercloud/v2/openstack/identity/v3/roles"
	_ "modernc.org/sqlite"
)

type SQLitePrincipalRepositories struct {
	db       *sql.DB
	Users    *SQLiteUserRepository
	Projects *SQLiteProjectRepository
	Roles    *SQLiteRoleRepository
}

func OpenSQLitePrincipalRepositories(
	dataSourceName string,
) (*SQLitePrincipalRepositories, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	repositories := &SQLitePrincipalRepositories{
		db:       db,
		Users:    &SQLiteUserRepository{db: db},
		Projects: &SQLiteProjectRepository{db: db},
		Roles:    &SQLiteRoleRepository{db: db},
	}
	if err := repositories.Init(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return repositories, nil
}

func (r *SQLitePrincipalRepositories) Init() error {
	statements := []string{
		`CREATE TABLE IF NOT EXISTS identity_users (
			sequence INTEGER PRIMARY KEY AUTOINCREMENT,
			id TEXT NOT NULL UNIQUE,
			name TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			default_project_id TEXT NOT NULL,
			description TEXT NOT NULL,
			domain_id TEXT NOT NULL,
			enabled INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS identity_projects (
			sequence INTEGER PRIMARY KEY AUTOINCREMENT,
			id TEXT NOT NULL UNIQUE,
			name TEXT NOT NULL UNIQUE,
			description TEXT NOT NULL,
			domain_id TEXT NOT NULL,
			enabled INTEGER NOT NULL,
			is_domain INTEGER NOT NULL
		)`,
		`CREATE TABLE IF NOT EXISTS identity_roles (
			sequence INTEGER PRIMARY KEY AUTOINCREMENT,
			id TEXT NOT NULL UNIQUE,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			domain_id TEXT NOT NULL
		)`,
	}

	for _, statement := range statements {
		if _, err := r.db.Exec(statement); err != nil {
			return err
		}
	}

	return nil
}

func (r *SQLitePrincipalRepositories) Close() error {
	return r.db.Close()
}

type SQLiteUserRepository struct {
	db *sql.DB
}

func (r *SQLiteUserRepository) Save(user appidentity.User) appidentity.User {
	_, err := r.db.Exec(
		`INSERT INTO identity_users (
			id,
			name,
			password,
			default_project_id,
			description,
			domain_id,
			enabled
		) VALUES (?, ?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			name = excluded.name,
			password = excluded.password,
			default_project_id = excluded.default_project_id,
			description = excluded.description,
			domain_id = excluded.domain_id,
			enabled = excluded.enabled`,
		user.ID,
		user.Name,
		user.Password,
		user.DefaultProjectID,
		user.Description,
		user.DomainID,
		boolToInt(user.Enabled),
	)
	if err != nil {
		panic(fmt.Errorf("save identity user: %w", err))
	}

	return user
}

func (r *SQLiteUserRepository) List() []appidentity.User {
	rows, err := r.db.Query(`
SELECT
	id,
	name,
	password,
	default_project_id,
	description,
	domain_id,
	enabled
FROM identity_users
ORDER BY sequence`)
	if err != nil {
		panic(fmt.Errorf("list identity users: %w", err))
	}
	defer rows.Close()

	users := []appidentity.User{}
	for rows.Next() {
		user, err := scanUser(rows)
		if err != nil {
			panic(fmt.Errorf("scan identity user: %w", err))
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Errorf("iterate identity users: %w", err))
	}

	return users
}

func (r *SQLiteUserRepository) Get(id string) (appidentity.User, error) {
	return r.getBy(`id = ?`, id)
}

func (r *SQLiteUserRepository) FindByName(
	name string,
) (appidentity.User, error) {
	return r.getBy(`name = ?`, name)
}

func (r *SQLiteUserRepository) getBy(
	condition string,
	value string,
) (appidentity.User, error) {
	row := r.db.QueryRow(`
SELECT
	id,
	name,
	password,
	default_project_id,
	description,
	domain_id,
	enabled
FROM identity_users
WHERE `+condition, value)

	user, err := scanUser(row)
	if errors.Is(err, sql.ErrNoRows) {
		return appidentity.User{}, appidentity.ErrUserNotFound
	}
	if err != nil {
		return appidentity.User{}, err
	}

	return user, nil
}

func (r *SQLiteUserRepository) Reset() {
	if _, err := r.db.Exec(`DELETE FROM identity_users`); err != nil {
		panic(fmt.Errorf("reset identity users: %w", err))
	}
}

type SQLiteProjectRepository struct {
	db *sql.DB
}

func (r *SQLiteProjectRepository) Save(project projects.Project) projects.Project {
	_, err := r.db.Exec(
		`INSERT INTO identity_projects (
			id,
			name,
			description,
			domain_id,
			enabled,
			is_domain
		) VALUES (?, ?, ?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			name = excluded.name,
			description = excluded.description,
			domain_id = excluded.domain_id,
			enabled = excluded.enabled,
			is_domain = excluded.is_domain`,
		project.ID,
		project.Name,
		project.Description,
		project.DomainID,
		boolToInt(project.Enabled),
		boolToInt(project.IsDomain),
	)
	if err != nil {
		panic(fmt.Errorf("save identity project: %w", err))
	}

	return project
}

func (r *SQLiteProjectRepository) List() []projects.Project {
	rows, err := r.db.Query(`
SELECT
	id,
	name,
	description,
	domain_id,
	enabled,
	is_domain
FROM identity_projects
ORDER BY sequence`)
	if err != nil {
		panic(fmt.Errorf("list identity projects: %w", err))
	}
	defer rows.Close()

	result := []projects.Project{}
	for rows.Next() {
		project, err := scanProject(rows)
		if err != nil {
			panic(fmt.Errorf("scan identity project: %w", err))
		}
		result = append(result, project)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Errorf("iterate identity projects: %w", err))
	}

	return result
}

func (r *SQLiteProjectRepository) Get(id string) (projects.Project, error) {
	return r.getBy(`id = ?`, id)
}

func (r *SQLiteProjectRepository) FindByName(
	name string,
) (projects.Project, error) {
	return r.getBy(`name = ?`, name)
}

func (r *SQLiteProjectRepository) getBy(
	condition string,
	value string,
) (projects.Project, error) {
	row := r.db.QueryRow(`
SELECT
	id,
	name,
	description,
	domain_id,
	enabled,
	is_domain
FROM identity_projects
WHERE `+condition, value)

	project, err := scanProject(row)
	if errors.Is(err, sql.ErrNoRows) {
		return projects.Project{}, appidentity.ErrProjectNotFound
	}
	if err != nil {
		return projects.Project{}, err
	}

	return project, nil
}

func (r *SQLiteProjectRepository) Reset() {
	if _, err := r.db.Exec(`DELETE FROM identity_projects`); err != nil {
		panic(fmt.Errorf("reset identity projects: %w", err))
	}
}

type SQLiteRoleRepository struct {
	db *sql.DB
}

func (r *SQLiteRoleRepository) Save(role roles.Role) roles.Role {
	_, err := r.db.Exec(
		`INSERT INTO identity_roles (
			id,
			name,
			description,
			domain_id
		) VALUES (?, ?, ?, ?)
		ON CONFLICT(id) DO UPDATE SET
			name = excluded.name,
			description = excluded.description,
			domain_id = excluded.domain_id`,
		role.ID,
		role.Name,
		role.Description,
		role.DomainID,
	)
	if err != nil {
		panic(fmt.Errorf("save identity role: %w", err))
	}

	return role
}

func (r *SQLiteRoleRepository) List() []roles.Role {
	rows, err := r.db.Query(`
SELECT
	id,
	name,
	description,
	domain_id
FROM identity_roles
ORDER BY sequence`)
	if err != nil {
		panic(fmt.Errorf("list identity roles: %w", err))
	}
	defer rows.Close()

	result := []roles.Role{}
	for rows.Next() {
		role, err := scanRole(rows)
		if err != nil {
			panic(fmt.Errorf("scan identity role: %w", err))
		}
		result = append(result, role)
	}
	if err := rows.Err(); err != nil {
		panic(fmt.Errorf("iterate identity roles: %w", err))
	}

	return result
}

func (r *SQLiteRoleRepository) Get(id string) (roles.Role, error) {
	row := r.db.QueryRow(`
SELECT
	id,
	name,
	description,
	domain_id
FROM identity_roles
WHERE id = ?`, id)

	role, err := scanRole(row)
	if errors.Is(err, sql.ErrNoRows) {
		return roles.Role{}, appidentity.ErrRoleNotFound
	}
	if err != nil {
		return roles.Role{}, err
	}

	return role, nil
}

func (r *SQLiteRoleRepository) Reset() {
	if _, err := r.db.Exec(`DELETE FROM identity_roles`); err != nil {
		panic(fmt.Errorf("reset identity roles: %w", err))
	}
}

type rowScanner interface {
	Scan(dest ...any) error
}

func scanUser(scanner rowScanner) (appidentity.User, error) {
	var user appidentity.User
	var enabled int
	err := scanner.Scan(
		&user.ID,
		&user.Name,
		&user.Password,
		&user.DefaultProjectID,
		&user.Description,
		&user.DomainID,
		&enabled,
	)
	if err != nil {
		return appidentity.User{}, err
	}
	user.Enabled = intToBool(enabled)

	return user, nil
}

func scanProject(scanner rowScanner) (projects.Project, error) {
	var project projects.Project
	var enabled int
	var isDomain int
	err := scanner.Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.DomainID,
		&enabled,
		&isDomain,
	)
	if err != nil {
		return projects.Project{}, err
	}
	project.Enabled = intToBool(enabled)
	project.IsDomain = intToBool(isDomain)
	project.Tags = []string{}
	project.Extra = map[string]any{}

	return project, nil
}

func scanRole(scanner rowScanner) (roles.Role, error) {
	var role roles.Role
	err := scanner.Scan(
		&role.ID,
		&role.Name,
		&role.Description,
		&role.DomainID,
	)
	if err != nil {
		return roles.Role{}, err
	}
	role.Links = map[string]any{}
	role.Extra = map[string]any{}
	role.Options = map[roles.Option]any{}

	return role, nil
}

func boolToInt(value bool) int {
	if value {
		return 1
	}

	return 0
}

func intToBool(value int) bool {
	return value != 0
}
