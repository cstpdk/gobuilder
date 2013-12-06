package main

type Project struct{
    Name string
    Descript string
    git string
    gitbranch string
    buildscript string
    buildkey string
}

var projectschema string =
`
CREATE TABLE project(
    name CHAR(100) PRIMARY KEY NOT NULL,
    descript TEXT,
    git CHAR(100),
    gitbranch CHAR(100),
    buildkey CHAR(64)
);
`
