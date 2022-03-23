#!/usr/bin/env bash

ORGPATH=$(pwd)
cd $(dirname $0)

. ./build.sh apinto-ingress-controller $1

buildApp
packageApp

cd ${ORGPATH}