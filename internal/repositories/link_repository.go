package repositories

import (
	"database/sql"
	"math/rand"
	"time"

	"demo/internal/models"

	"github.com/google/uuid"
)

type LinkRepository struct {
	db *sql.DB
}

func NewLinkRepository(db *sql.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (r *LinkRepository) Save(e *models.Link) (*models.Link, error) {
	e.ID = uuid.New()
	e.Code = GeneratorCode(10)
	e.CreateAt = time.Now()
	stmt, err := r.db.Prepare(`
		INSERT INTO links (id, code, url, create_at)
		VALUE (?, ?, ?, ?)
	`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(e.ID, e.Code, e.Url, e.CreateAt)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	return e, nil
}

func (r *LinkRepository) FindAll() ([]*models.Link, error) {
	stmt, err := r.db.Prepare(`
		SELECT * FROM links
	`)

	if err != nil {
		return nil, err
	}

	var links []*models.Link

	rows, err := stmt.Query()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var l models.Link
		err := rows.Scan(&l.ID, &l.Code, &l.Url, &l.CreateAt)

		if err != nil {
			return nil, err
		}

		links = append(links, &l)
	}

	// var ids []uuid.UUID

	// rows, err := stmt.Query()

	// if err != nil {
	// 	return nil, err
	// }

	// for rows.Next() {
	// 	var i uuid.UUID
	// 	err := rows.Scan(&i)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	ids = append(ids, i)
	// }

	// if len(ids) <= 0 {
	// 	return nil, nil
	// }

	// var links []*models.Link

	// for _, id := range ids {
	// 	l, err := GetLink(id, r.db)

	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	links = append(links, l)
	// }

	defer stmt.Close()
	defer rows.Close()

	return links, nil
}

func (r *LinkRepository) FindByID(id uuid.UUID) (*models.Link, error) {
	return GetLink(id, r.db)
}

func GetLink(id uuid.UUID, db *sql.DB) (*models.Link, error) {
	stmt, err := db.Prepare(`
		SELECT * FROM Links where id = ?
	`)

	if err != nil {
		return nil, err
	}

	var l models.Link

	rows, err := stmt.Query(id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&l.ID, &l.Code, &l.Url, &l.CreateAt)
		if err != nil {
			return nil, err
		}
	}

	defer stmt.Close()
	defer rows.Close()

	return &l, nil
}

func GeneratorCode(length int) string {
	src_rand := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src_rand)
	var src = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	result := make([]byte, length)

	for i := range result {
		result[i] = src[r.Intn(len(src))]
	}

	return string(result)
}

func GetLinkByCode(code string, db *sql.DB) (*models.Link, error) {
	stmt, err := db.Prepare(`
		SELECT * FROM Links where code = ?
	`)

	if err != nil {
		return nil, err
	}

	var l models.Link

	rows, err := stmt.Query(code)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&l.ID, &l.Code, &l.Url, &l.CreateAt)
		if err != nil {
			return nil, err
		}
	}

	defer stmt.Close()
	defer rows.Close()

	return &l, nil
}

func (r *LinkRepository) FindByCode(code string) (*models.Link, error) {
	return GetLinkByCode(code, r.db)
}
