package main

import(
    "errors"
    "code.google.com/p/go.crypto/bcrypt"
)

/*
User a user
*/
type User struct{
    Username string `db:"username" json:"username"`
    Email string `db:"email" json:"email"`
    Role string `db:"role" json:"role"`
}

/*
Loginuser the same as User but the password, used for authentication purposes
*/
type Loginuser struct{
    Username string `db:"username" json:"username"`
    Password string `db:"password" json:"password"`
    Email string `db:"email" json:"email"`
    Role string `db:"role" json:"role"`
}

//Validates the role
func rolevalidation(r string) bool{
    return (r == "admin" || r == "user")
}

/*
Hashpwd hashes the users password using bcrypt
*/
func (u *Loginuser) Hashpwd(){
    b := []byte(u.Password)
    r, err := bcrypt.GenerateFromPassword(b, 12)
    if err != nil {
        panic(err)
    }
    u.Password = string(r)
}

var userschema =
`
CREATE TABLE user(
    username CHAR(256) PRIMARY KEY NOT NULL,
    password text NOT NULL,
    email text,
    role text
);
`

/*
CreateUser creates a new user in the database
*/
func CreateUser(u Loginuser) (User, error){
    //Validate role
    if !rolevalidation(u.Role) {
        return User{}, errors.New("invalid role")
    }

    u.Hashpwd()

    _, err := db.NamedExec(
        `INSERT INTO user (username, password, email, role)
        VALUES(:username, :password, :email, :role)`, u)

    if err != nil{
        return User{}, errors.New("user already exists")
    }

    return User{u.Username, u.Email, u.Role}, nil
}

/*
UpdateUser updates the user in the database
*/
func UpdateUser(u Loginuser) (User, error){

    if !rolevalidation(u.Role){
        return User{}, errors.New("invalid role")
    }

    u.Hashpwd()

    _, err := db.NamedExec(
        `UPDATE user SET password=:password, email=:email,
        role=:role WHERE username=:username`, u)

    if err != nil{
        return User{}, errors.New("could not update user")
    }

    return User{u.Username, u.Email, u.Role}, nil
}

/*
DeleteUser deletes the user from the database
*/
func DeleteUser(name string) error{
    _, err := db.Exec(`DELETE FROM user WHERE username=$1`, name)

    return err
}

/*
GetUsers gets the user from the database
*/
func GetUser(name string) (User, error){
    u := User{}

    err := db.Get(&u, `SELECT username, email, role 
    FROM user WHERE username=$1`, name)

    if err != nil {
        return User{}, err
    }

    return u, nil
}

/*
Get users
*/
func GetUsers() []User {
    users := []User{}
    db.Select(&users, "SELECT username, email, role FROM user")

    return users
}

//Get a loginuser
func getloginuser(name string) (Loginuser, error){
    u := Loginuser{}

    err := db.Get(&u, `SELECT username, password, email, role 
    FROM user WHERE username=$1`, name)

    if err != nil {
        return Loginuser{}, err
    }

    return u, nil
}
