package repository

import (
	"test-go-api/config"
	"test-go-api/entity"
)

func GetDataUser() ([]*entity.User, error) {
	db, err := config.MySQLConnection()

	if err != nil {
		return nil, err
	}

	rows, errRows := db.Query("SELECT id, username, email FROM users")
	if errRows != nil {
		return nil, errRows
	}

	var result []*entity.User

	for rows.Next() {
		var user = entity.User{}
		errScan := rows.Scan(&user.ID, &user.Username, &user.Email)
		if errScan != nil {
			return nil, errScan
		}
		result = append(result, &user)
	}

	return result, nil
}
