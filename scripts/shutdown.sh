#!/bin/bash

#获取项目pid后，直接杀死项目

BROWSER_BIN="project_bin"
pid=`ps -ef | grep ${BROWSER_BIN} | grep -v grep | awk '{print $2}'`
if [ ! -z ${pid} ];then
    kill -9 $pid
fi
