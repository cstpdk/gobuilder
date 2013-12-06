package main

import(
    "github.com/codegangsta/martini"
    "net/http"
)

type User struct{
    Username string
    Email string
    Role string
}

type Loginuser struct{
    Username string
    Password string
    Email string
    Role string
}

/*
Auth middel ware function for handling authentication and injecting User
*/
func Auth(c martini.Context, w http.ResponseWriter, r *http.Request){
    var u User = User{
        Username:  "test",
        Email: "test",
        Role:  "user",
    }
    c.Map(u)
}

/*
Admin middle ware function for handling admin check user after Auth
*/
func Admin(u User, w http.ResponseWriter, r *http.Request){
    if u.Role != "admin"{
        http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
    }
}
