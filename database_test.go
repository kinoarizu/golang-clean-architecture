package belajar_golang_database

import (
	"database/sql"
	"testing"
	_ "github.com/lib/pq"
)

func TestOpenConnection(t *testing.T){
	db, err := sql.Open("postgres", "postgres://postgres:pass123@localhost:5432/belajar_postgres")
	if err != nil {
		panic(err)
	}
	defer db.Close()
}