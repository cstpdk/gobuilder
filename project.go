package main

import(
    "errors"
    "os"
    "path/filepath"
    "fmt"
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

/*
ProjectFolder gets the project folder's path
*/
func (p *Project) ProjectFolder() string {
    return filepath.Join(workspace,p.Name)
}

/*
CreateProjectFolder create the project folder
*/
func (p *Project) CreateProjectFolder() {
    err := os.MkdirAll(p.ProjectFolder(), 0700)

    if err != nil {
        panic(err)
    }
}

/*
BuildScriptFile get the project's buildscript file
*/
func (p *Project) BuildScriptFile() *os.File {
    path := filepath.Join(p.ProjectFolder(), "buildscript.sh")

    f, err := os.Create(path)

    if err != nil {
        panic(err)
    }

    return f
}

/*
WriteBuildScript writes the buildscript to the buildscript file
*/
func (p *Project) WriteBuildScript() {
    _, err := fmt.Fprint(p.BuildScriptFile(), p.Buildscript)

    if err != nil {
        panic(err)
    }
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
        return Project{}, errors.New("project already exists")
    }

    p.CreateProjectFolder()

    p.WriteBuildScript()

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
UpdateProject updates the project in the database
*/
func UpdateProject(p Project) (Project, error){

    _, err := db.NamedExec(`UPDATE project SET description=:description,
    git=:git, gitbranch=:gitbranch, buildkey=:buildkey WHERE name=:name`, p)

    if err != nil{
        return Project{}, errors.New("could not update project")
    }

    //TODO: Update file with buildscript

    return p, nil
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
