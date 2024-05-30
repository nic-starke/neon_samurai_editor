#!/usr/bin/env bash

BUILD_DIR="build/dev"
APP_NAME="neosam_ed"

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR" || exit

# Clean
echo "Cleaning.."
rm -f ./${BUILD_DIR}/**/${APP_NAME}*
mkdir -p ${BUILD_DIR}

# Build
echo "Building..."
go build -o ${BUILD_DIR}/${APP_NAME} . && echo "Build completed"
