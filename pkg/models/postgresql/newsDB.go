package postgresql

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
	"news-project/pkg/models"
)

type NewsModel struct {
	DB *sql.DB
}

func (m *NewsModel) Insert(title, content, category string) (int, error) {
	stmt := `INSERT INTO news (title, content, date, category) VALUES($1, $2, CURRENT_TIMESTAMP, $3) RETURNING id`
	var id int
	err := m.DB.QueryRow(stmt, title, content, category).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (m *NewsModel) Get(id int) (*models.News, error) {
	stmt := `SELECT id, title, content, date, category FROM news WHERE id = $1`
	row := m.DB.QueryRow(stmt, id)
	n := &models.News{}
	err := row.Scan(&n.ID, &n.Title, &n.Content, &n.Date, &n.Category)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrorMessage
		} else {
			return nil, err
		}
	}

	return n, nil
}

func (m *NewsModel) Latest() ([]*models.News, error) {
	stmt := `SELECT id, title, content, date, category FROM news ORDER BY date DESC LIMIT 10`
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	newsList := []*models.News{}
	for rows.Next() {
		n := &models.News{}
		err = rows.Scan(&n.ID, &n.Title, &n.Content, &n.Date, &n.Category)
		if err != nil {
			return nil, err
		}
		newsList = append(newsList, n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return newsList, nil
}

func (m *NewsModel) GetByCategory(category string) ([]*models.News, error) {
	stmt := `
        SELECT id, title, content, category, date FROM news
        WHERE category = $1
        ORDER BY date DESC`
	rows, err := m.DB.Query(stmt, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	newsArray := make([]*models.News, 0)
	for rows.Next() {
		n := &models.News{}
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &n.Category, &n.Date)
		if err != nil {
			return nil, err
		}
		newsArray = append(newsArray, n)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return newsArray, nil
}

func (m *NewsModel) DeleteNewsById(id int) error {
	query := `
		DELETE FROM news WHERE id = $1`

	stmt, err := m.DB.Prepare(query)

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
