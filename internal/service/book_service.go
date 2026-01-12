package service

import (
	"api-rest-go/internal/model"
	"api-rest-go/internal/store"
	"errors"
)

type Logger interface {
	Log(msg, error string)
}

type Service struct {
	store  store.Store
	logger Logger
}

func New(s store.Store) *Service {
	return &Service{store: s}
}

func (s *Service) ObtenerTodosLibros() ([]*model.Libro, error) {
	if s.logger != nil {
		s.logger.Log("info", "estamos obteniendo los libros")
	}
	libros, err := s.store.GetAll()
	if err != nil {
		if s.logger != nil {
			s.logger.Log("error", "error al obtener los libros: "+err.Error())
		}
		return nil, err
	}
	return libros, nil
}

func (s *Service) ObtenerLibroPorID(id int) (*model.Libro, error) {
	return s.store.GetByID(id)
}

func (s *Service) CrearLibro(libro *model.Libro) (*model.Libro, error) {
	if libro.Titulo == "" {
		return nil, errors.New("El título del libro no puede estar vacío")
	}
	return s.store.Create(libro)
}

func (s *Service) ActualizarLibro(id int, libro *model.Libro) (*model.Libro, error) {
	return s.store.Update(id, libro)
}

func (s *Service) EliminarLibro(id int) error {
	return s.store.Delete(id)
}
