package main

import(
    "github.com/codegangsta/martini"
    "net/http"
)

func Route(m *martini.ClassicMartini){
    m.Get("/", func(res http.ResponseWriter, req *http.Request) {
        http.ServeFile(res, req, "API.md")
    })
}
