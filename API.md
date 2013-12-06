#API

##Project

|Method | URL                            |Input                      |Output                       | Description                  |  
|-------|--------------------------------|---------------------------|-----------------------------|------------------------------|
|GET    | /projects                      | params: search, page      | \[project,project..\]       | Gets all projects            |
|GET    | /project/:name                 |                           | project                     | Get project with id          | 
|GET    | /project/:name/buildhistory    | params: page              | \[build,build..\]           | The build history            |
|POST   | /project                       | project                   | project                     | Create a new project         |
|POST   | /project/:name/build           | params: key               | build                       | Build the project            |
|PUT    | /project                       | project                   | project                     | Update a project             |
|DELETE | /project                       | project                   | bool                        | Delete a project             |

##User

|Method | URL                            |Input                      |Output                       | Description                  |  
|-------|--------------------------------|---------------------------|-----------------------------|------------------------------|
|GET    | /users                         | params: search, page      | \[user, user..\]            | Gets all users               |
|GET    | /user/:name                    |                           | user                        | Gets a users info            |
|POST   | /user                          | loginuser                 | user                        | Create a new user (admin)    |
|PUT    | /user                          | loginuser                 | user                        | Updated a user (admin/self)  |
|DELETE | /user                          | user                      | bool                        | Delete a user (admin)        |
                             
##Build
|Method | URL                            |Input                      |Output                       | Description                  |  
|-------|--------------------------------|---------------------------|-----------------------------|------------------------------|
|GET    | /builds                        | params: page              | \[build,build..\]           | Get current running builds   |
|GET    | /build/:id                     |                           | build                       | Get info on the build        |
|GET    | /build/:id/log                 |                           | buildlog                    | Get the build log            |

##Authentication
Using HTTP basic authentication

* Authorization: Basic base64(username:password)
* See more at [Wikipedia](http://en.wikipedia.org/wiki/Basic_access_authentication#Client_side)

For building without access to the system (remote or via other system) a key can
be provided in the URL:

/project/:id/build?key={key}



#JSON
##project:
```json
{
    "name"          : "A unique name",
    "description"   : "A description of the project",
    "git"           : "github.com/user/repo",
    "gitbranch"     : "master",
    "buildscript"   : "build script",
    "buildkey"      : "key to build project"
}
```

##user
```json
{
    "username" : "username",
    "email"    : "email",
    "role"     : "{admin, user}" 
}
```

##loginuser
```json
{
    "username" : "username",
    "password" : "password",
    "email"    : "email",
    "role"     : "{admin, user}" 
}
```


##build
```json
{
    "id"        : "id",
    "user"      : "user that have started build",
    "complete"  : true
}
```

##build log
First and last entry is used to get partial build logs on the next request.
```json
{
    "buildid"       : "build id",
    "log"           : "log output",
    "firstentry"    : "id of entry",
    "lastentry"     : "id of entry"
}
```
