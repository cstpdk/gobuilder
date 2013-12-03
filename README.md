#gobuilder
A build server written in Go

#API

##GET
| URI                       | Description                  |  
|---------------------------|------------------------------|
| /projects                 | Gets all projects            |
| /project/:id              | Get project with id          | 
| /project/:id/history      | The build history            |
| /project/:id/status       | Get build status             | 
| /builds                   | Get current running builds   |

##POST
| URI                       | Description                  |  
|---------------------------|------------------------------|
| /project/:id/build        | Build the project            |

##PUT

##DELETE

#JSON
##project:
```json
{
    "id"   : "uuid",
    "name" : "A name",
    "description" : "A description of the project",
    "git"  : "github.com/user/repo",
    "gitbranch" : "master",
    "buildlinux" : "build script for linux",
    "buildwin"   : "build script for windows"
}
```
