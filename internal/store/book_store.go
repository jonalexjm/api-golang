package store


import (
	"../model"
	"database/sql"
)

type Store interface {
	GetAll()([]*model.Libro, error)
	GetByID(id int)(*model.Libro, error)
	Create(book *model.Libro)(*model.Libro, error)
	Update(id int, book *model.Libro)(*model.Libro, error)
	Delete(id int)(error)
}

type store struct {
   db *sql.DB
}

func NewStore(db *sql.DB) Store {
   return &store{db: db}
}

func (s *store) GetAll()([]*model.Libro, error) {
	q := s.db.Query("SELECT id, title, author FROM books")

	rows, err : = s.db.Query(q)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libros []*model.Libro
	for rows.Next() {
		var libro model.Libro
		if err := rows.Scan(&libro.ID, &libro.Title, &libro.Author); err != nil {
			return nil, err
		}
		libros = append(libros, &libro)
	}
	return libros, nil
}
func (s *store) GetByID(id int)(*model.Libro, error) {
	q := `SELECT id, title, author FROM books WHERE id = ?`
	var libro *model.Libro
	row := s.db.QueryRow(q, id).Scan(&libro.ID, &libro.Title, &libro.Author); 
	if err != nil {
		return nil, err
	}
	return libro, nil
}

func (s *store) Create(libro *model.Libro)(*model.Libro, error) {
	q := `INSERT INTO books (title, author) VALUES (?, ?)`
	result, err := s.db.Exec(q, libro.Title, libro.Author)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	libro.ID = int(id)
	return libro, nil
}

func (s *store) Update(id int, libro *model.Libro)(*model.Libro, error) {
	q := `UPDATE books SET title = ?, author = ? WHERE id = ?`
	_, err := s.db.Exec(q, libro.Title, libro.Author, id);
	if err != nil {
		return nil, err
	}

	libro.ID = id
	return libro, nil
}

func (s *store) Delete(id int)(error) {
	q := `DELETE FROM books WHERE id = ?`
	_, err := s.db.Exec(q, id);
	if err != nil {
		return err
	}
	return nil
}