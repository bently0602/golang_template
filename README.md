# golang template

Tired of GOPATH? Want a project based path environment for GO? Need your dev environment to be cross platform? Want easy cross compiling? This could help you. It's my typical GO dev environment.

Uses [go-task](https://github.com/go-task/task) for task configuration and [uuid](https://github.com/uutils/coreutils) for Windows build of coreutils.

## How to Use

- Download ZIP of repo and extract to where you want your dev project.
- Change name of folder in src to reflect the package name you want for your app.
- Change Taskvars.yml / BINARY_NAME: to be the name of that package.
- Change Taskvars.yml / CONFIG_NAME: to be the name of your config file (no path, config file is expected to be in root).
- Change Taskvars.yml / DEPS: to add any urls of dependancies you want (By GO GET).
- Change LICENSE to represent your license.
- Change README.md as your readme markdown file.
- Run: *build.sh(or .bat) deps*
- Run: *build.sh*
- Run: *build.sh run*

## Structure

* /bin - Contains current dev build by running just "build.sh". "build.sh run" runs this. __build.sh creates this.__
* /static - Folder is copied over into releases when built out. Meant to hold static www files or files need by the application. __build.sh creates this if you do not.__
* /releases - Folder gets built and contains packaged releases for each platform built for. __build.sh creates this.__
* /src - Source code folder. Should contain all deps and your package folder.
* /utils - Contains built copys of go-task and uuid.
* config.json - Template of a config file for your application. CONFIG_NAME in Taskvars.yml specifies the name build.sh looks for.
* README.md - This file. Modify it to be what you want.

## Build.sh(.bat) Options

build.bat/build.sh {all, run, deps, clean, fmt}

1. __build.sh__ (no params)

   builds project for current machine, then immediately executes it if there are no errors in the build process

2. __build.sh run__

   runs latest built project (with path of the golang_template root)

3. __build.sh deps__

   grabs all dependencies in Taskvars.yml

4. __build.sh clean__

   removes bin and releases folder requiring a rebuild

5. __build.sh fmt__

   go fmt over the package

6. __build.sh all__

   builds and packages in releases folder. copies over in each release folder: static folder, config.json (or whatever CONFIG_NAME is), and README.md.
   
   *amd64:*
   
    * os x
    * windows
    * linux
    * openbsd
    * plan 9
   
   *arm:*
   
    * linux
    
   *Modify Taskfile.yml to specify different release platforms.*

## Issues

Windows problems delete golang deps from src folder:

`
del /f /s /q golang.org 1>nul
rmdir /s /q golang.org
`