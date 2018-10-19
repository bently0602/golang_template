# golang template

Tired of GOPATH? Want a project based path environment for GO? This could help you. It's my typical GO dev environment.

## Using

- download ZIP of repo.
- change name of folder in src to reflect the package name you want for your app.
- change Taskvars.yml / BINARY_NAME: to be the name of that package and add any urls of dependancies you want.
- run: build.sh deps
- run: build.sh
- run: build.sh run

_Windows: Use cygwin_

build.sh {all, run, deps, clean, fmt}

1. build.sh (no params)
   builds project for current machine
2. build.sh run
    - copies config.json from src folder to root of project folder
    - runs latest built project
3. build.sh deps
   grabs all dependencies in Taskvars.yml
4. build.sh clean
   removes bin and releases folder
5. build.sh fmt
   go fmt over the package
6. build.sh all
   builds amd64 releases in releases folder for:
    * os x
    * windows
    * linux
    * openbsd
   modify Taskfile.yml to specify different release platforms
