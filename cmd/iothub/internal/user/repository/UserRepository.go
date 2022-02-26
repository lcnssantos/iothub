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

func (this UserRepository) Create(data dto.CreateUserDto, ctx context.Context) error {
	return database.ExecuteTransaction(ctx, this.database, func(tx *sql.Tx) error {

		prepare, err := tx.Prepare("INSERT INTO users (name, email, password, active) values ($1, $2, $3, false)")

		if err != nil {
			return err
		}

		_, err = prepare.Exec(data.Name, data.Email, data.Password)

		return err
	})
}

func (this UserRepository) FindOneByEmail(email string, ctx context.Context) (*dto.User, error) {
	prepare, err := this.database.PrepareContext(ctx, "SELECT * FROM users where email = $1 limit 1")

	if err != nil {
		return nil, err
	}

	user, err := this.scanEntity(prepare.QueryRow(email))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this UserRepository) FindOneById(id uint64, ctx context.Context) (*dto.User, error) {
	prepare, err := this.database.PrepareContext(ctx, "SELECT * FROM users WHERE id = $1 LIMIT 1")

	if err != nil {
		return nil, err
	}

	user, err := this.scanEntity(prepare.QueryRow(id))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this UserRepository) List(ctx context.Context) ([]*dto.User, error) {
	prepare, err := this.database.PrepareContext(ctx, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	rows, err := prepare.Query()

	if err != nil {
		return nil, err
	}

	return this.scanEntities(rows)
}

func (this UserRepository) scanEntities(r *sql.Rows) ([]*dto.User, error) {
	users := make([]*dto.User, 0)

	for r.Next() {
		user := new(dto.User)
		err := r.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (this UserRepository) scanEntity(r *sql.Row) (*dto.User, error) {
	var user = dto.User{}
	err := r.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}
