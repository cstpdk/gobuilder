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

    //Post a new user
    m.Post("/user",  Auth, Admin, binding.Json(Loginuser{}), binding.ErrorHandler, func(u Loginuser, j Json) (int,string){
        user, err := CreateUser(u)

        if err != nil{
            return http.StatusConflict, j(err.Error())
        }

        return http.StatusOK, j(user)
    })

    //Put an existing user
    m.Put("/user", Auth, binding.Json(Loginuser{}), binding.ErrorHandler,
    func(u Loginuser, user User, j Json) (int, string){
        if u.Username != user.Username && user.Role != "admin" {
            return http.StatusUnauthorized, j("Access denied")
        }

        r, err := UpdateUser(u)

        if err != nil {
            return http.StatusConflict, j(err.Error())
        }

        return http.StatusOK, j(r)
    })
}
