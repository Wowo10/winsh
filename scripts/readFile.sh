#!/bin/sh

if [ -z "$1" ]
then
    echo 'Give me a file!'
else
    cat $1
fi