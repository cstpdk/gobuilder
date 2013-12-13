package main

type Build struct{
    id  int
    user string
    project string
    complete bool
}

var buildschema string =
`
CREATE TABLE build(
    id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
    user CHAR(256),
    project CHAR(256),
    complete INT DEFAULT 0,
    FOREIGN KEY(user) REFERENCES user(username),
    FOREIGN KEY(project) REFERENCES project(name)
);
`
