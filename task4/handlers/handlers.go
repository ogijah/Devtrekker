package handlers

import (
	"devtrekker/data"
	"encoding/json"
	"io"
	"mime"
	"net/http"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

const path = "data/telephones.json"

func (h *Handler) GetTelephones(rw http.ResponseWriter, req *http.Request) {
	output, err := data.GetTelephones(path)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	renderJSON(rw, output)
}

func (h *Handler) GetTelephoneById(rw http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(rw, "bad parameter of id", http.StatusBadRequest)
		return
	}
	input, _ := data.GetTelephoneById(id, path)
	if (input == data.Input{}) {
		http.Error(rw, "there is no telephone number by that id", http.StatusNotFound)
		return
	}
	renderJSON(rw, input)
}

func (h *Handler) DeleteTelephone(rw http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(rw, "bad parameter of id", http.StatusBadRequest)
		return
	}
	input, _ := data.GetTelephoneById(id, path)
	if (input == data.Input{}) {
		http.Error(rw, "there is no telephone number by that id", http.StatusNotFound)
		return
	}
	data.DeleteTelephone(id, path)
}

func (h *Handler) UploadTelephone(rw http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	if mediatype != "application/json" {
		http.Error(rw, "expect application/json Content-Type", http.StatusUnsupportedMediaType)
		return
	}

	newInput, err := decodeBody(req.Body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	validate := validator.New()
	err = validate.Struct(newInput)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = data.UploadTelephone(*newInput, path)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
}

func decodeBody(r io.Reader) (*data.Input, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()
	var newInput data.Input
	if err := dec.Decode(&newInput); err != nil {
		return nil, err
	}
	return &newInput, nil
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(js)
	if err != nil {
		return
	}
}
