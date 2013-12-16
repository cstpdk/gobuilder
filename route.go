package main

import(
    "github.com/codegangsta/martini"
    "github.com/codegangsta/martini-contrib/binding"
    "net/http"
    "errors"
)

/*
Route assigns the martini routes
*/
func Route(m martini.Router){
    m.Get("/", func(res http.ResponseWriter, req *http.Request) {
        http.ServeFile(res, req, "API.md")
    })
    userroutes(m)
}

/*
userroutes assigns the user related routes
*/
func userroutes(m martini.Router){
    //Post a new user
    m.Post("/user",  Auth, Admin, binding.Json(Loginuser{}),
    binding.ErrorHandler, func(u Loginuser) (int,interface{}){
        user, err := CreateUser(u)

        if err != nil{
            return http.StatusConflict, err
        }

        return http.StatusCreated, user
    })

    //Put an existing user
    m.Put("/user", Auth, binding.Json(Loginuser{}), binding.ErrorHandler,
    func(u Loginuser, user User) (int, interface{}){
        if u.Username != user.Username && user.Role != "admin" {
            return http.StatusUnauthorized, errors.New("Unauthorized")
        }

        r, err := UpdateUser(u)

        if err != nil {
            return http.StatusConflict, err
        }

        return http.StatusOK, r
    })

    //Delete a user
    m.Delete("/user/:name", Auth, func(user User,
    params martini.Params) (int, interface{}){
        name := params["name"]
        if user.Username != name && user.Role != "admin" {
            return http.StatusUnauthorized, errors.New("Access denied")
        }

        err := DeleteUser(name)

        if err != nil {
            return http.StatusConflict, err
        }

        return http.StatusOK, ""
    })

    m.Get("/user/:name", Auth, func(params martini.Params)(int,
    interface{}){
        name := params["name"]
        u, err := GetUser(name)

        if err != nil {
            return http.StatusNotFound, err
        }

        return http.StatusOK, u
    })

    m.Get("/users", Auth, func() []User{
        return GetUsers()
    })
}
