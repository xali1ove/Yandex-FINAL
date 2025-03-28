package handler

import dbfuncs "github.com/xali1ove/Yandex-FINAL/database"

type Handler struct {
	DB *dbfuncs.DB
}

func NewHandler(db *dbfuncs.DB) *Handler {
	return &Handler{DB: db}
}
