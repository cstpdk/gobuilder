package main

import(
    _ "github.com/mattn/go-sqlite3"
    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

/*
SetupDB sets up the database
The database is using sqlite and connects to the specified file
*/
func SetupDB(dbfile string){
    //Setup database connection
    db = sqlx.MustConnect("sqlite3", dbfile)
    db.Execf(userschema)
    db.Execf(projectschema)
    db.Execf(buildschema)
}
