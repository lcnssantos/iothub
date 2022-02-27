package repository

import (
	"context"
	"database/sql"

	"github.com/lcnssantos/iothub/cmd/iothub/internal/user/dto"
	"github.com/lcnssantos/iothub/internal/database"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database: database}
}

func (r UserRepository) GetTransaction(ctx context.Context) (*sql.Tx, error) {
	return database.GetTransaction(ctx, r.database)
}

func (r UserRepository) Create(data dto.CreateUserDto, tx *sql.Tx) error {
	prepare, err := tx.Prepare("INSERT INTO users (name, email, password, active) values (?, ?, ?, false)")

	if err != nil {
		return err
	}

	_, err = prepare.Exec(data.Name, data.Email, data.Password)

	return err

}

func (r UserRepository) FindOneByEmail(email string, ctx context.Context) (*dto.User, error) {
	prepare, err := r.database.PrepareContext(ctx, "SELECT * FROM users where email = ? limit 1")

	if err != nil {
		return nil, err
	}

	user, err := r.scanEntity(prepare.QueryRow(email))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) FindOneById(id uint64, ctx context.Context) (*dto.User, error) {
	prepare, err := r.database.PrepareContext(ctx, "SELECT * FROM users WHERE id = ? LIMIT 1")

	if err != nil {
		return nil, err
	}

	user, err := r.scanEntity(prepare.QueryRow(id))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r UserRepository) List(ctx context.Context) ([]*dto.User, error) {
	prepare, err := r.database.PrepareContext(ctx, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	rows, err := prepare.Query()

	if err != nil {
		return nil, err
	}

	return r.scanEntities(rows)
}

func (r UserRepository) scanEntities(row *sql.Rows) ([]*dto.User, error) {
	users := make([]*dto.User, 0)

	for row.Next() {
		user := new(dto.User)
		err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r UserRepository) scanEntity(row *sql.Row) (*dto.User, error) {
	var user = dto.User{}
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}
