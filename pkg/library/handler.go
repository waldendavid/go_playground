package library

import "net/http"

func NewHandler(service Library) Handler {
	return &handler{service: service}
}

type Handler interface {
	GetBooks() http.HandlerFunc
}

type handler struct {
	service Library
}

func (h *handler) GetBooks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// ... implementacja json marshalingu itd
		books, err := h.service.GetBooks(r.Context())
		if err != nil {
			// TODO zwrócić HTTp internal error
		}
		//TODO zwrócić JSON marshall Books, czyli booksy do JSONa
	}
}
