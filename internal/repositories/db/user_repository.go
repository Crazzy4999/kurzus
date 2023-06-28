package db

import (
	"database/sql"
	"errors"
	"fmt"
	"hangryAPI/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Create(u *models.User) error {
	stmt, err := ur.db.Prepare("INSERT INTO users (first_name, last_name, email, password) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return errors.New("couldn't prepare statment to insert new user")
	}

	user, _ := ur.GetUserByEmail(u.Email)
	if user != nil {
		return errors.New("user already exist with this email")
	}

	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.Password)
	if err != nil {
		fmt.Println(err)
		return errors.New("insert into users failed")
	}

	return nil
}

func (ur *UserRepository) GetAll() ([]*models.User, error) {
	stmt, err := ur.db.Prepare("SELECT id, address_id, first_name, last_name, email, password FROM users")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to get all users")
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, errors.New("getting all users failed")
	}
	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		err = rows.Scan(&user.ID, &user.Address, &user.FirstName, &user.LastName, &user.Email, &user.Password)
		if err == sql.ErrNoRows {
			return nil, nil
		} else if err != nil {
			return nil, errors.New("types mismatch during the scanning")
		}
		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	stmt, err := ur.db.Prepare("SELECT id, address_id, first_name, last_name, email, password FROM users WHERE users.email = $1")
	if err != nil {
		return nil, errors.New("couldn't prepare statement to get user by email")
	}

	var user models.User

	row := stmt.QueryRow(email)
	err = row.Scan(&user.ID, &user.Address, &user.FirstName, &user.LastName, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) Update(u *models.User) error {
	stmt, err := ur.db.Prepare("UPDATE users SET address_id = $2, first_name = $1, last_name = $3, password = $4 WHERE users.id = $5")
	if err != nil {
		return errors.New("couldn't prepare statement for updating user")
	}

	_, err = stmt.Exec(u.Address, u.FirstName, u.LastName, u.Password, u.ID)
	if err != nil {
		return errors.New("updating user failed")
	}

	return nil
}

func (ur *UserRepository) Delete(id int) error {
	stmt, err := ur.db.Prepare("DELETE FROM users WHERE users.id = $1")
	if err != nil {
		return errors.New("couldn't prepare statement for deleting user by id")
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return errors.New("deleting user by id failed")
	}

	return nil
}
