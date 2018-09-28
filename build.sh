#!/bin/sh

if grep -q Microsoft /proc/version; then
    # this does not work atm
    #   windows path cant append form wsl
    #   without more work
    export PATH=$PATH:/mnt/c/Go/bin
    export GOROOT=/mnt/c/Go
    ./utils/task/task_windows_amd64/task $@
else
    case "$(uname -s)" in

    Darwin)
        ./utils/task/task_darwin_amd64/task $@
        ;;

    Linux)
        ./utils/task/task_linux_amd64/task $@
        ;;

    # these carry over and translate the win path so GO should
    # be available from a standard windows install
    CYGWIN*|MINGW64*|MSYS*)
        ./utils/task/task_windows_amd64/task $@
        ;;

    *)
        echo 'other not handled OS, see build.sh script' 
        ;;
    esac
fi
