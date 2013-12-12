package main

import(
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/binding"
    "net/http"
)

func Route(m martini.Router){
    m.Get("/", func(res http.ResponseWriter, req *http.Request) {
        http.ServeFile(res, req, "API.md")
    })
    m.Post("/user",  Auth, Admin, binding.Json(Loginuser{}), binding.ErrorHandler, func(u Loginuser, j Json) (int,string){
        user, err := CreateUser(u)

        if err != nil{
            return 409, j(err.Error())
        }

        return 200, j(user)
    })
}
