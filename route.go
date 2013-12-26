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
    projectroutes(m)
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

    //Get a specific user
    m.Get("/user/:name", Auth, func(params martini.Params)(int,
    interface{}){
        name := params["name"]
        u, err := GetUser(name)

        if err != nil {
            return http.StatusNotFound, errors.New("could not find user")
        }

        return http.StatusOK, u
    })

    //Get all users
    m.Get("/users", Auth, GetUsers)
}

/*
projectroutes assigns the project related routes
*/
func projectroutes(m martini.Router){
    //Post a new project
    m.Post("/project", Auth, binding.Json(Project{}), binding.ErrorHandler,
    func(p Project) (int, interface{}){
        project, err := CreateProject(p)
        if err != nil {
            return http.StatusConflict, err
        }
        return http.StatusOK, project
    })

    //Get a specific user
    m.Get("/project/:name", Auth, func(params martini.Params)(int,
    interface{}){
        name := params["name"]
        p, err := GetProject(name)

        if err != nil {
            return http.StatusNotFound, errors.New("could not find project")
        }

        return http.StatusOK, p
    })

    //Get all projects
    m.Get("/projects", Auth, GetProjects)
}
