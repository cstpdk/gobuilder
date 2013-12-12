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
    UserRoutes(m)
}

func UserRoutes(m martini.Router){
    //Post a new user
    m.Post("/user",  Auth, Admin, binding.Json(Loginuser{}),
    binding.ErrorHandler, func(u Loginuser, j Json) (int,string){
        user, err := CreateUser(u)

        if err != nil{
            return http.StatusConflict, j(err)
        }

        return http.StatusCreated, j(user)
    })

    //Put an existing user
    m.Put("/user", Auth, binding.Json(Loginuser{}), binding.ErrorHandler,
    func(u Loginuser, user User, j Json) (int, string){
        if u.Username != user.Username && user.Role != "admin" {
            return http.StatusUnauthorized, j("Access denied")
        }

        r, err := UpdateUser(u)

        if err != nil {
            return http.StatusConflict, j(err)
        }

        return http.StatusOK, j(r)
    })

    //Delete a user
    m.Delete("/user/:name", Auth, func(user User, j Json,
    params martini.Params) (int, string){
        name := params["name"]
        if user.Username != name && user.Role != "admin" {
            return http.StatusUnauthorized, j("Access denied")
        }

        err := DeleteUser(name)

        if err != nil {
            return http.StatusConflict, j("Could not delete user")
        }

        return http.StatusOK, ""
    })

    m.Get("/user/:name", Auth, func(j Json, params martini.Params)(int,
    string){
        name := params["name"]
        u, err := GetUser(name)

        if err != nil {
            return http.StatusNotFound, j("Could not find user")
        }

        return http.StatusOK, j(u)
    })

    m.Get("/users", Auth, func(j Json) (int, string){
        users := GetUsers()
        return http.StatusOK, j(users)
    })
}
