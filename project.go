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
    name text PRIMARY KEY NOT NULL,
    descript text,
    git text,
    gitbranch text,
    buildkey text
);
`
