package document

import (
	"database/sql"
	"fmt"
	"net/http"
)

type session struct {
	sql.DB
}

type State string

type DocumentType string

type Document struct {
	ID      int
	Name    string
	DocType DocumentType
	State   State
}

type DocumentHandler interface {
	GetDocument(http.ResponseWriter, http.Request)
	UpdateDocument(http.ResponseWriter, http.Request)
	DeleteDocument(http.ResponseWriter, http.Request)
}

func (d Document) GetDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Print("unimplemented")
}

func (d Document) UpdateDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Print("unimplemented")
}

func (d Document) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	fmt.Print("unimplemented")
}
