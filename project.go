package main

type Project struct{
    Name string `db:"name" json:"name"`
    Description string `db:"description" json:"description"`
    git string `db:"git" json:"git"`
    gitbranch string `db:"gitbranch" json:"gitbranch"`
    buildscript string `db:"buildscript" json:"buildscript"`
    buildkey string `db:"buildkey" json:"buildkey"`
}

var projectschema string =
`
CREATE TABLE project(
    name CHAR(256) PRIMARY KEY NOT NULL,
    description text,
    git text,
    gitbranch text,
    buildkey text
);
`
