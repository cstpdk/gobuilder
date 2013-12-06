package main

type Build struct{
    id  int
    user string
    complete bool
}

var buildschema string =
`
CREATE TABLE build(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    user text,
    complete INT DEFAULT 0
);
`
