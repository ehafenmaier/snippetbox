package models

import (
	"database/sql"
	"errors"
	"time"
)

// Snippet type holds the data for an individual snippet
type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// SnippetModel type wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// Insert a new snippet into the database.
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
             VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// Get a specific snippet based on its id
func (m *SnippetModel) Get(id int) (Snippet, error) {
	// Create the SQL query
	stmt := `SELECT id, title, content, created, expires FROM snippets
             WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// Execute query
	row := m.DB.QueryRow(stmt, id)

	// Initialize a new zeroed Snippet struct.
	var s Snippet

	// Copy returned data into the snippet struct
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Snippet{}, ErrNoRecord
		}

		return Snippet{}, err
	}

	return s, nil
}

// Latest returns the 10 most recently created snippets
func (m *SnippetModel) Latest() ([]Snippet, error) {
	// Create the SQL query
	stmt := `SELECT id, title, content, created, expires FROM snippets
             WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// Execute the query
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Initialize a slice for snippets
	var snippets []Snippet

	// Loop through the results and copy into the snippets slice
	for rows.Next() {
		var s Snippet

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}

		snippets = append(snippets, s)
	}

	// Check for any iteration errors
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}
