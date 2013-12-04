#gobuilder
A build server written in Go

##API

###GET
| URL                            |Input                      |Output                       | Description                  |  
|--------------------------------|---------------------------|-----------------------------|------------------------------|
| /projects                      |                           |                             | Gets all projects            |
| /project/:id                   |                           |                             | Get project with id          | 
| /project/:id/buildhistory      |                           |                             | The build history            |
| /builds                        |                           |                             | Get current running builds   |
| /build/:id                     |                           |                             | Get info on the build        |
| /build/:id/log                 |                           |                             | Get info on the build        |

###POST
| URL                            |Input                      |Output                       | Description                  |  
|--------------------------------|---------------------------|-----------------------------|------------------------------|
| /project/:id/build             |None                       |build                        | Build the project            |
                             
###PUT                       
                             
###DELETE                    

###Authentication
Using HTTP basic authentication

* Authorization: Basic base64(username:password)
* See more at [Wikipedia](http://en.wikipedia.org/wiki/Basic_access_authentication#Client_side)

For building without access to the system (remote or via other system) a key can
be provided in the URL:

/project/:id/build?key={key}



##JSON
###project:
```json
{
    "id"   : "uuid",
    "name" : "A name",
    "description" : "A description of the project",
    "git"  : "github.com/user/repo",
    "gitbranch" : "master",
    "buildlinux" : "build script for linux",
    "buildwin"   : "build script for windows",
    "buildkey"   : "key to build project"
}
```

###user
```json
{
    "username" : "username",
    "email"    : "email",
    "
}
```

###build
```json
{
    "id" : "id",
    "user" : "user that have started build",
    "complete" : true
}
```

###build log
First and last entry is used to get partial build logs on the next request.
```json
{
    "buildid" : "build id",
    "log" : "log output",
    "firstentry" : "id of entry",
    "lastentry" : "id of entry"
}
```
