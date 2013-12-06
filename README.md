#Gobuilder
A build server written in Go

##Dependencies

* [Martini](https://github.com/codegangsta/martini) for the web service.
* [sqlx](https://github.com/jmoiron/sqlx) for the database.

##API

###GET

| URL                            |Input                      |Output                       | Description                  |  
|--------------------------------|---------------------------|-----------------------------|------------------------------|
| /projects                      | params: search, page      | \[project,project..\]       | Gets all projects            |
| /project/:name                 |                           | project                     | Get project with id          | 
| /project/:name/buildhistory    | params: page              | \[build,build..\]           | The build history            |
| /builds                        | params: page              | \[build,build..\]           | Get current running builds   |
| /build/:id                     |                           | build                       | Get info on the build        |
| /build/:id/log                 |                           | buildlog                    | Get the build log            |
| /users                         | params: search, page      | \[user, user..\]            | Gets all users               |
| /user/:name                    |                           | user                        | Gets a users info            |

###POST

| URL                            |Input                      |Output                       | Description                  |  
|--------------------------------|---------------------------|-----------------------------|------------------------------|
| /project                       | project                   | project                     | Create a new project         |
| /user                          | user                      | user                        | Create a new user            |
| /project/:name/build           | params: key               | build                       | Build the project            |
                             
###PUT                       

| URL                            |Input                      |Output                       | Description                  |  
|--------------------------------|---------------------------|-----------------------------|------------------------------|
| /project                       | project                   | project                     | Create a new project         |
| /user                          | user                      | user                        | Create a new user (admin)    |
                             
###DELETE                    

| URL                            |Input                      |Output                       | Description                  |  
|--------------------------------|---------------------------|-----------------------------|------------------------------|
| /project                       | project                   | bool                        | Delete a project             |
| /user                          | user                      | bool                        | Delete a user (admin)        |

###Authentication
Using HTTP basic authentication

* Authorization: Basic base64(username:password)
* See more at [Wikipedia](http://en.wikipedia.org/wiki/Basic_access_authentication#Client_side)

For building without access to the system (remote or via other system) a key can
be provided in the URL:

/project/:id/build?key={key}



