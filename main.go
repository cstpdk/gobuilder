package main

import(
    "github.com/codegangsta/martini"
    "os"
    "flag"
    _ "github.com/mattn/go-sqlite3"
    "github.com/jmoiron/sqlx"
)

var db *sqlx.DB
func main(){
    //Flags
    port := flag.String("port","3000","Port to host on")

    flag.Parse()

    //Setup martini
    m := martini.New()
    m.Use(martini.Recovery())
    m.Use(martini.Logger())

    os.Setenv("PORT", *port)
    r := martini.NewRouter()

    //Setup routes
    Route(r)

    //Setup database connection
    db = sqlx.MustConnect("sqlite3","database.db")

    //Start server
    m.Action(r.Handle)
    m.Run()
}
