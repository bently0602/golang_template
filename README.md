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
- Run: "build.sh(or .bat) deps"
- Run: "build.sh"
- Run: "build.sh run"

## Structure

* /bin - Contains current dev build by running just "build.sh". "build.sh run" runs this.
* /static - Folder is copied over into releases when built out. Meant to hold static www files or files need by the application.
* /releases - Folder gets built and contains packaged releases for each platform built for.
* /src - Source code folder. Should contain all deps and your package folder.
* /utils - Contains built copys of go-task and uuid.
* config.json

## Build.sh(.bat) Options

build.bat/build.sh {all, run, deps, clean, fmt}

1. build.sh (no params)
   builds project for current machine

2. build.sh run
   runs latest built project (with path of the golang_template root)

3. build.sh deps
   grabs all dependencies in Taskvars.yml

4. build.sh clean
   removes bin and releases folder requiring a rebuild

5. build.sh fmt
   go fmt over the package

6. build.sh all
   builds and package in releases folder
   amd64:
    * os x
    * windows
    * linux
    * openbsd
    * plan 9
   arm:
    * linux
   modify Taskfile.yml to specify different release platforms
