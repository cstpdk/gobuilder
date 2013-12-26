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

/*
CreateProject creates a new project in the database
*/
func CreateProject(p Project) (Project, error){
    _, err := db.NamedExec(`
    INSERT INTO project (name, description, git, gitbranch, buildkey)
    VALUES(:name, :description, :git, :gitbranch, :buildkey)`, p)

    if err != nil {
        return Project{}, err
    }

    //TODO: Make workspace dir and insert buildscript in file

    return p, nil
}

/*
GetProject get specific project from the database.
*/
func GetProject(name string) (Project, error) {
    p := Project{}

    err := db.Get(&p,
    `SELECT name, description, git, gitbranch, buildkey FROM project WHERE
    name=$1`, name)

    //TODO: Get information from files

    return p, err
}

/*
DeleteProject deletes the project from the database
*/
func DeleteProject(name string) error{
    _, err := db.Exec(`DELETE FROM project WHERE name=$1`, name)

    //TODO: Delete workspace folder

    return err
}

/*
GetProjects gets all the projects in the database.
Projects does not contain buildscript
*/
func GetProjects() []Project {
    projects := []Project{}
    db.Select(&projects,
    `SELECT name, description, git, gitbranch, buildkey FROM project`)

    return projects
}
