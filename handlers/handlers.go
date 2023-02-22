package handlers

import (
	"github.com/ogijah/Devtrekker/data"
	"log"
	"net/http"
)

type Handler struct {
	l *log.Logger
}

func (p *Handler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	data.GetTelephones()
}

func NewHandler(l *log.Logger) *Handler {
	return &Handler{l}
}
