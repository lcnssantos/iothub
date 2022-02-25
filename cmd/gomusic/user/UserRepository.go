package user

import "database/sql"

type Repository struct {
	database *sql.DB
}

func NewRepository(database *sql.DB) *Repository {
	return &Repository{database: database}
}

func (this Repository) Create(data CreateUserDto) error {
	prepare, err := this.database.Prepare("INSERT INTO users (name, email, password, active) values ($1, $2, $3, false)")

	if err != nil {
		return err
	}

	_, err = prepare.Exec(data.Name, data.Email, data.Password)

	return err
}

func (this Repository) FindOneByEmail(email string) (*User, error) {
	prepare, err := this.database.Prepare("SELECT * FROM users where email = $1 limit 1")

	if err != nil {
		return nil, err
	}

	user, err := this.scanEntity(prepare.QueryRow(email))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this Repository) FindOneById(uid string) (*User, error) {
	prepare, err := this.database.Prepare("SELECT * FROM users WHERE id = $1 LIMIT 1")

	if err != nil {
		return nil, err
	}

	user, err := this.scanEntity(prepare.QueryRow(uid))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this Repository) scanEntity(r *sql.Row) (*User, error) {
	var user = User{}
	err := r.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.Active, &user.CreatedAt, &user.UpdatedAt)
	return &user, err
}
