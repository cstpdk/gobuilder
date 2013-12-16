package main

import(
    "github.com/codegangsta/martini"
    "net/http"
    "encoding/base64"
    "strings"
    "code.google.com/p/go.crypto/bcrypt"
)

/*
Auth middel ware function for handling authentication and injecting User
*/
func Auth(c martini.Context, w http.ResponseWriter, r *http.Request){
    auth := r.Header.Get("Authorization")
    split := strings.SplitN(string(auth)," ",2)

    if len(split) < 2 {
        http.Error(w, "Unauthorized", 401)
    }

    data, err := base64.StdEncoding.DecodeString(split[1])


    if err != nil {
        http.Error(w, err.Error(), 500)
    }

    info := strings.SplitN(string(data),":",2)

    if len(info) < 2 {
        http.Error(w, "Unauthorized", 401)
    }

    u, _ := getloginuser(info[0])

    crypterr := bcrypt.CompareHashAndPassword([]byte(u.Password),
    []byte(info[1]))

    if crypterr != nil{
        http.Error(w, "Unauthorized", 401)
    }

    user := User{ Username : u.Username, Email : u.Email, Role : u.Role }
    c.Map(user)
}

/*
Admin middle ware function for handling admin check user after Auth
*/
func Admin(u User, w http.ResponseWriter, r *http.Request){
    if u.Role != "admin"{
        http.Error(w, http.StatusText(http.StatusForbidden),
        http.StatusForbidden)
    }
}
