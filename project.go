package main

import(
)

/*
Project a project
*/
type Project struct{
    Name string `db:"name" json:"name"`
    Description string `db:"description" json:"description"`
    Git string `db:"git" json:"git"`
    Gitbranch string `db:"gitbranch" json:"gitbranch"`
    Buildscript string `db:"buildscript" json:"buildscript"`
    Buildkey string `db:"buildkey" json:"buildkey"`
}

var projectschema =
`
CREATE TABLE project(
    name CHAR(256) PRIMARY KEY NOT NULL,
    description text,
    git text,
    gitbranch text,
    buildkey text
);
`

func CreateProject(p Project) (Project, error){
    _, err := db.NamedExec(`
    INSERT INTO project (name, description, git, gitbranch, buildkey)
    VALUES(:name, :description, :git, :gitbranch, :buildkey)`, p)

    //TODO: Make workspace dir and insert buildscript in file

    if err != nil {
        return Project{}, err
    }

    return p, nil
}
