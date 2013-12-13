package main

/*
Build a build of a project
*/
type Build struct{
    Id  int
    User string
    Project string
    Complete bool
}

var buildschema =
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
