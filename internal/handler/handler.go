package handler

import "github.com/MamushevArup/stat4market/internal/repository"

type Handler struct {
	repository *repository.Repository
}

func New(r *repository.Repository) *Handler {
	return &Handler{repository: r}
}
