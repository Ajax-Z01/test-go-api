package repository

import (
	"database/sql"

	"test-go-api/entity"

	_ "github.com/go-sql-driver/mysql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAllUsers() ([]*entity.User, error) {
	users := make([]*entity.User, 0)
	query := "SELECT id, username, email FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(&user.ID, &user.USERNAME, &user.EMAIL)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) GetUserByID(id int) (*entity.User, error) {
	user := &entity.User{}
	query := "SELECT id, username, email FROM users WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.USERNAME, &user.EMAIL)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) CreateUser(user *entity.User) error {
	query := "INSERT INTO users (username, email) VALUES (?, ?)"
	_, err := r.db.Exec(query, user.USERNAME, user.EMAIL)
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUser(user *entity.User) error {
	query := "UPDATE users SET username = ?, email = ? WHERE id = ?"
	_, err := r.db.Exec(query, user.USERNAME, user.EMAIL, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUser(id int) error {
	stmt, err := ur.db.Prepare("DELETE FROM users WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
