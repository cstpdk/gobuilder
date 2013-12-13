package main

import(
    "github.com/codegangsta/martini"
    "os"
    "flag"
)

var workspace string

func main(){
    //Flags
    port := flag.String("port","3000","Port to host on")
    database := flag.String("database","database.db","Database file to use, creates one if it does not exist")
    workspc := flag.String("workspace","workspace/","Location of the workspace to use")

    flag.Parse()

    workspace = *workspc

    //Setup martini
    m := martini.New()
    m.Use(martini.Recovery())
    m.Use(martini.Logger())
    m.Use(JSONEncoder)

    os.Setenv("PORT", *port)
    r := martini.NewRouter()

    //Setup routes
    Route(r)

    //Setup the database and connection
    SetupDB(*database)


    //Start server
    m.Action(r.Handle)
    m.Run()
}
