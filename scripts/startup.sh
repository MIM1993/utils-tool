#!/bin/bash

BROWSER_BIN="project_bin"

go build -o ${BROWSER_BIN} ../src
echo "Success build "
CONFIG_PATH="../configs/"
nohup ./${BROWSER_BIN} -config ${CONFIG_PATH} >output 2>&1 &
