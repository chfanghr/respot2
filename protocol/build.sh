#!/usr/bin/env bash

set -u
stty -echo

goPath=${GOPATH}

msg() {
    echo "${1}:${2-}"
}

panic() {
    msg "build" "failed"
    exit 1
}

install_protoc() {
    if [[ $(command uname) == "Linux" ]]
    then
        command sudo apt-get -y install protobuf-compiler > /dev/null 2>&1
        return $?
    fi
    if [[ $(command uname) == "Darwin" ]]
    then
        command brew install protobuf > /dev/null 2>&1
        return $?
    fi
    msg "install_protoc" "unsupported platform"
    return 1
}

check_protoc() {
    if command -v protoc > /dev/null 2>&1
    then
        msg "check_protoc" "command found"
        return 0
    else
        msg "check_protoc" "command not found,install it"
        install_protoc
        if [[ "$?" -ne 0 ]]
        then
            msg "check_protoc" "installation failed"
            return 1
        fi
    fi
}

install_protoc_gen_go() {
    command go get -u -v github.com/golang/protobuf/protoc-gen-go > /dev/null 2>&1
    return $?
}

check_protoc_gen_go() {
    if [[ -x "$goPath/bin/protoc-gen-go" ]]
    then
        msg "check_protoc_gen_go" "command found"
        return 0
    else
        msg "check_protoc_gen_go" "command not found,install it"
        install_protoc_gen_go
        if [[ "$?" -ne 0 ]]
        then
            msg "check_protoc" "installation failed"
            return 1
        fi
    fi
}

check_go() {
    if command -v go > /dev/null 2>&1
    then
        msg "check_go" "command found"
        return 0
    else
        msg "check_go" "command not found"
        return 1
    fi
}

check_dependencies() {
    msg "check_dependencies" "begin"
    check_go
    if [[ "$?" -ne 0 ]]; then
        msg "check_dependencies" "failed"
        return 1
    fi
    check_protoc
    if [[ "$?" -ne 0 ]]; then
        msg "check_dependencies" "failed"
        return 1
    fi
    check_protoc_gen_go
    if [[ "$?" -ne 0 ]]
    then
        msg "check_dependencies" "failed"
        return 1
    fi
    msg "check_dependencies" "done"
    return 0
}

build_protocol() {
    msg "build_protocol" "begin"
    if [[ ! -d "spotify" ]]
    then
        command mkdir spotify
    fi
    command protoc -I proto/ --plugin=${goPath}/bin/protoc-gen-go --go_out spotify/ proto/*.proto > /dev/null 2>&1
    if [[ "$?" -ne 0 ]]
    then
        msg "build_protocol" "failed"
        return 1
    fi
    local version=$(command git rev-parse --short HEAD)
    echo -n > spotify/spotify.go
    echo "package spotify" >> spotify/spotify.go
    echo >> spotify/spotify.go
    echo "var Version = \"v$version\" " >> spotify/spotify.go
    go fmt spotify/*.go >> spotify/spotify.go
    msg "build_protocol" "done"
    return 0
}

main() {
    if [[ ! "$goPath" ]]
    then
        $goPath=${HOME}/go
    fi
    msg "build" "begin"
    check_dependencies
    if [[ "$?" -ne 0 ]]
    then
        panic
    fi
    build_protocol
    if [[ "$?" -ne 0 ]]
    then
        panic
    fi
    msg "build" "done"
    exit 0
}

clean() {
    command rm -rf spotify > /dev/null 2>&1
}

if [[ "$#" -eq "1" ]]
then
    if [[ "$1" == "clean" ]]
    then
        clean
    fi
else
    main
fi
