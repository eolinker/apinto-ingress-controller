#!/usr/bin/env bash

function genVersion() {
    if [[ $1 = "" ]]
    then
      v=$(git rev-parse --short HEAD)
      time=$(date "+%Y%m%d%H")
      echo "$time-$v"
      exit 0
    fi
    echo "$1"
}

function buildApp(){
    cd ${BasePath}
    rm -rf ${OutputPath}
    mkdir -p ${OutputPath}

    buildCMD="go build  -o ${OutputPath}/$APP"
    echo "build $APP:${buildCMD}"
    ${buildCMD}

    if [[ "$?" != "0" ]]
    then
        rm -rf $OutputPath
        exit 1
    fi
    echo "$Version" > ${OutputPath}/version

}
function packageApp() {
    cd ${OutputPath}
    tar -zcf "${OutputPath}.linux.x64.tar.gz" --xform "s#^#${APP}/#"  *
    cd "${BasePath}"
}


cd ../..
BasePath=$(pwd)
Version=$(genVersion $2)
APP=$1
echo $APP
echo $Version
OutputName="${APP}-${Version}"
OutputPath="${BasePath}/out/${OutputName}"