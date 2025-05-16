package repositories

import (
	"database/sql"
	"demo/internal/models"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func GetUser(id uuid.UUID, db *sql.DB) (*models.User, error) {
	stmt, err := db.Prepare(`
		SELECT * FROM users WHERE id = ?
	`)

	if err != nil {
		return nil, err
	}

	var user models.User

	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Actived, &user.CreateAt, &user.UpdateAt)
		if err != nil {
			return nil, err
		}
	}

	defer stmt.Close()
	defer rows.Close()

	return &user, nil
}

func (ur *UserRepository) FindAll() ([]*models.User, error) {
	stmt, err := ur.db.Prepare(`
		SELECT id FROM users
	`)

	if err != nil {
		return nil, err
	}

	var ids []uuid.UUID

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id uuid.UUID
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}

	var users []*models.User

	for _, id := range ids {
		user, err := GetUser(id, ur.db)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	defer stmt.Close()
	defer rows.Close()

	return users, nil
}

func (ur *UserRepository) FindByID(id uuid.UUID) (*models.User, error) {
	return GetUser(id, ur.db)
}

func (ur *UserRepository) Exists(id uuid.UUID) (bool, error) {
	stmt, err := ur.db.Prepare(`
		SELECT id FROM users WHERE id = ?
	`)

	if err != nil {
		return false, err
	}

	rows, err := stmt.Query(id)

	if err != nil {
		return false, err
	}

	defer stmt.Close()
	defer rows.Close()

	return rows.Next(), nil
}

func (ur *UserRepository) Save(e *models.User) (*models.User, error) {
	exists, err := ur.Exists(e.ID)
	if err != nil {
		return nil, err
	}

	if exists {
		user, err := GetUser(e.ID, ur.db)
		if err != nil {
			return nil, err
		}

		user.Username = e.Username
		user.Password = e.Password
		user.Email = e.Email
		user.Actived = e.Actived
		user.UpdateAt = time.Now()

		stmt, err := ur.db.Prepare(`
			UPDATE users SET username = ?, password = ?, email = ?, actived = ?, update_at = ? WHERE id = ?
		`)

		if err != nil {
			return nil, err
		}

		_, err = stmt.Exec(user.Username, user.Password, user.Email, user.Actived, user.UpdateAt, user.ID)
		if err != nil {
			return nil, err
		}

		defer stmt.Close()

		return user, nil
	}

	e.ID = uuid.New()
	e.Actived = true
	e.CreateAt = time.Now()
	e.UpdateAt = time.Now()
	stmt, err := ur.db.Prepare(`
		INSERT INTO users (id, username, password, email, actived, create_at, update_at)
		VALUE (?, ?, ?, ?, ?, ?, ?)
	`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(e.ID, e.Username, e.Password, e.Email, e.Actived, e.CreateAt, e.UpdateAt)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return e, nil
}
