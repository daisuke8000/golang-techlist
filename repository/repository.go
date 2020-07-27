package repository

import(
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

//setDB
func SetDB(d *sqlx.DB)  {
	db = d
}