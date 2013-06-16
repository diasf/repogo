#!/bin/bash

CURDIR=`pwd`
SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd $SCRIPTDIR
revel run github.com/diasf/repogo

cd $CURDIR
