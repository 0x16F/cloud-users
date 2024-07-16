package users

import (
	"context"
	"strings"

	"github.com/0x16F/cloud-users/internal/entity"
	"github.com/huandu/go-sqlbuilder"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pkg/errors"
)

const (
	limit = 1000
)

type Repo struct {
	db *pgxpool.Conn
}

func NewRepo(db *pgxpool.Conn) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	query := `
		INSERT INTO cd_users (email, username, password, salt)
		VALUES (@email, @username, @password, @salt)
		RETURNING id
	`

	args := pgx.NamedArgs{
		"email":    user.Email,
		"username": user.Username,
		"password": user.Password,
		"salt":     user.Salt,
	}

	var id uint64

	if err := r.db.QueryRow(ctx, query, args).Scan(&id); err != nil {
		return entity.User{}, errors.Wrap(err, "failed to create user")
	}

	user.ID = id

	return user, nil
}

func (r *Repo) GetUser(ctx context.Context, id uint64) (entity.User, error) {
	query := `
		SELECT id, email, username, password, salt
		FROM cd_users
		WHERE id = @id
	`

	args := pgx.NamedArgs{
		"id": id,
	}

	var user entity.User

	err := r.db.QueryRow(ctx, query, args).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Salt)
	if err != nil {
		return entity.User{}, errors.Wrap(err, "failed to get user")
	}

	return user, nil
}

func (r *Repo) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	query := `
		SELECT id, email, username, password, salt
		FROM cd_users
		WHERE email = LOWER(@email)
	`

	args := pgx.NamedArgs{
		"email": email,
	}

	var user entity.User

	err := r.db.QueryRow(ctx, query, args).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Salt)
	if err != nil {
		return entity.User{}, errors.Wrap(err, "failed to get user by email")
	}

	return user, nil
}

func (r *Repo) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	query := `
		SELECT id, email, username, password, salt
		FROM cd_users
		WHERE username = LOWER(@username)
	`

	args := pgx.NamedArgs{
		"username": username,
	}

	var user entity.User

	err := r.db.QueryRow(ctx, query, args).Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Salt)
	if err != nil {
		return entity.User{}, errors.Wrap(err, "failed to get user by username")
	}

	return user, nil

}

func (r *Repo) GetUsers(ctx context.Context, params entity.GetUsersParams) ([]entity.User, error) {
	if params.Limit == 0 {
		params.Limit = limit
	}

	sb := sqlbuilder.NewSelectBuilder()

	sb.Select("id", "email", "username", "password", "salt")
	sb.From("cd_users")
	sb.Limit(params.Limit)

	if params.LastID != 0 {
		sb.Where(sb.GT("id", params.LastID))
	}

	if params.Username != "" {
		sb.Where(sb.Like("username", strings.ToLower(params.Username)))
	}

	if params.Email != "" {
		sb.Where(sb.Like("email", strings.ToLower(params.Email)))
	}

	query, args := sb.Build()

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get users")
	}

	users := []entity.User{}

	for rows.Next() {
		var user entity.User

		if err := rows.Scan(&user.ID, &user.Email, &user.Username, &user.Password, &user.Salt); err != nil {
			return nil, errors.Wrap(err, "failed to scan user")
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *Repo) UpdateEmail(ctx context.Context, id uint64, email string) error {
	query := `
		UPDATE cd_users
		SET email = @email
		WHERE id = @id
	`

	args := pgx.NamedArgs{
		"id":    id,
		"email": email,
	}

	if _, err := r.db.Exec(ctx, query, args); err != nil {
		return errors.Wrap(err, "failed to update email")
	}

	return nil
}

func (r *Repo) UpdateUsername(ctx context.Context, id uint64, username string) error {
	query := `
		UPDATE cd_users
		SET username = @username
		WHERE id = @id
	`

	args := pgx.NamedArgs{
		"id":       id,
		"username": username,
	}

	if _, err := r.db.Exec(ctx, query, args); err != nil {
		return errors.Wrap(err, "failed to update username")
	}

	return nil
}

func (r *Repo) UpdatePassword(ctx context.Context, id uint64, password string, salt string) error {
	query := `
		UPDATE cd_users
		SET password = @password, salt = @salt
		WHERE id = @id
	`

	args := pgx.NamedArgs{
		"id":       id,
		"password": password,
		"salt":     salt,
	}

	if _, err := r.db.Exec(ctx, query, args); err != nil {
		return errors.Wrap(err, "failed to update password")
	}

	return nil
}

func (r *Repo) DeleteUser(ctx context.Context, id uint64) error {
	query := `
		UPDATE cd_users
		SET deleted_at = NOW()
		WHERE id = @id
	`

	args := pgx.NamedArgs{
		"id": id,
	}

	if _, err := r.db.Exec(ctx, query, args); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}
