package main

import(
    "github.com/codegangsta/martini"
    "os"
    "flag"
)
func main(){
    //Flags
    port := flag.String("port","3000","Port to host on")

    flag.Parse()

    //Setup martini
    os.Setenv("PORT", *port)
    m := martini.Classic()

    //Setup routes
    Route(m)

    //Start server
    m.Run()
}
