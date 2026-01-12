package transport

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"api-rest-go/internal/model"
	"api-rest-go/internal/service"
)

type BookHandler struct {
	service *service.Service
}

func New(s *service.Service) *BookHandler {
	return &BookHandler{service: s}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		libros, err := h.service.ObtenerTodosLibros()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libros)

	case http.MethodPost:
		var libro model.Libro
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		created, err := h.service.CrearLibro(&libro)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(created)
	// Otros métodos HTTP (POST, PUT, DELETE) se manejarían aquí
	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}

func (h *BookHandler) HandleBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/books/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		libro, err := h.service.ObtenerLibroPorID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(libro)

	case http.MethodPut:
		var libro model.Libro
		if err := json.NewDecoder(r.Body).Decode(&libro); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updated, err := h.service.ActualizarLibro(id, &libro)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(updated)

	case http.MethodDelete:
		err := h.service.EliminarLibro(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}
}
