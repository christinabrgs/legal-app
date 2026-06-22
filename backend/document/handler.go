package document

import (
	"database/sql"
	"net/http"
)

type session struct {
	sql.DB
}

type Document struct {
	ID    int
	Name  string
	Type  string
	State string
}

func (d Document) getDocument(w http.ResponseWriter, r *http.Request) {
}

func (d Document) updateDocument(w http.ResponseWriter, r *http.Request) {
}

func (d Document) deleteDocument(w http.ResponseWriter, r *http.Request) {
}
