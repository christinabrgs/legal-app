

create table document (
  id serial PRIMARY KEY
  name string
  doc_type string
  state string
)

create table state (
  state string PRIMARY KEY
)

create table document_type (
  doc_type string PRIMARY KEY
)
