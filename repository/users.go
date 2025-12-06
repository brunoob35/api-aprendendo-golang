package repository

import (
	"Api-Aula1/models"
	"database/sql"
)

type UsersRepo struct {
	db *sql.DB
}

func NewUsersRepo(db *sql.DB) *UsersRepo {
	return &UsersRepo{db}
}

func (u UsersRepo) Create(user models.Users) (int8, error) {
	query := `INSERT INTO treehousedb.users(
                            name,
                            email,
                            password,
                            cpf)
                            VALUES (?,?,?,?)`

	statement, err := u.db.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Password, user.CPF)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int8(uint64(lastid)), nil
}

func (u UsersRepo) FetchByEmail(email string) (DBuser models.Users, err error) {
	line, err := u.db.Query(`SELECT id, email, password FROM treehousedb.users WHERE email = ?`, email)
	if err != nil {
		return models.Users{}, err
	}
	defer line.Close()

	if line.Next() {
		if err := line.Scan(&DBuser.ID, &DBuser.Email, &DBuser.Password); err != nil {
			return models.Users{}, err
		}
	}

	return DBuser, nil
}

func (u UsersRepo) FetchByID(id uint64) (models.Users, error) {
	line, err := u.db.Query(`
        SELECT id, name, email, password, cpf
        FROM treehousedb.users
        WHERE id = ?`,
		id,
	)
	if err != nil {
		return models.Users{}, err
	}
	defer line.Close()

	var DBuser models.Users

	if line.Next() {
		if err := line.Scan(
			&DBuser.ID,
			&DBuser.Name,
			&DBuser.Email,
			&DBuser.Password,
			&DBuser.CPF,
		); err != nil {
			return models.Users{}, err
		}
	}

	// se não achou, volta struct vazia
	return DBuser, nil
}
