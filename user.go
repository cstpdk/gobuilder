package main

import(
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

var userschema string =
`
CREATE TABLE user(
    username text PRIMARY KEY NOT NULL,
    password text NOT NULL,
    email text,
    role text
);
`
