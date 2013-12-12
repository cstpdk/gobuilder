package main

import(
    "errors"
)

type User struct{
    Username string `db:"username" json:"username"`
    Email string `db:"email" json:"email"`
    Role string `db:"role" json:"role"`
}

type Loginuser struct{
    Username string `db:"username" json:"username"`
    Password string `db:"password" json:"password"`
    Email string `db:"email" json:"email"`
    Role string `db:"role" json:"role"`
}

func RoleValidation(r string) bool{
    return (r == "admin" || r == "user")
}

var userschema string =
`
CREATE TABLE user(
    username text PRIMARY KEY NOT NULL,
    password text NOT NULL,
    email text,
    role text
);
`

func CreateUser(u Loginuser) (User, error){
    //Validate role
    if !RoleValidation(u.Role) {
        return User{}, errors.New("Invalid role")
    }
    _, err := db.NamedExec(
        `INSERT INTO user (username, password, email, role)
        VALUES(:username, :password, :email, :role)`, u)

    if err != nil{
        return User{}, errors.New("User already exists")
    }

    return User{u.Username, u.Email, u.Role}, nil
}
