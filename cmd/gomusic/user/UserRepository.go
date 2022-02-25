package user

import (
	"context"
	"database/sql"

	"github.com/lcnssantos/gomusic/internal/database"
)

type Repository struct {
	database *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{database: database}
}

func (this Repository) Create(data CreateUserDto, ctx context.Context) error {
	return database.ExecuteTransaction(ctx, this.database, func(tx *sql.Tx) error {

		prepare, err := tx.Prepare("INSERT INTO users (name, email, password, active) values ($1, $2, $3, false)")

		if err != nil {
			return err
		}

		_, err = prepare.Exec(data.Name, data.Email, data.Password)

		return err
	})
}

func (this Repository) FindOneByEmail(email string, ctx context.Context) (*User, error) {
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

func (this Repository) FindOneById(uid string, ctx context.Context) (*User, error) {
	prepare, err := this.database.PrepareContext(ctx, "SELECT * FROM users WHERE id = $1 LIMIT 1")

	if err != nil {
		return nil, err
	}

	user, err := this.scanEntity(prepare.QueryRow(uid))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this Repository) List(ctx context.Context) ([]*User, error) {
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

func (this Repository) scanEntities(r *sql.Rows) ([]*User, error) {
	users := make([]*User, 0)

	for r.Next() {
		user := new(User)
		err := r.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (this Repository) scanEntity(r *sql.Row) (*User, error) {
	var user = User{}
	err := r.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}
