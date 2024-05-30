#!/usr/bin/env bash

# Constants
BUILD_DIR="build/release"
APP_NAME="neosam_ed"

# Build flags
FLAG_ARCH=${FLAG_ARCH:-"386 amd64 arm64"}
FLAG_OS=${FLAG_OS:-linux darwin windows}
FLAG_EXCLUDE_OSARCH="!darwin/arm !darwin/386"
VER_FLAG=$(git describe --always --long --dirty)
LD_FLAGS="-s -w"

# Get the parent directory of where this script is.
SOURCE="${BASH_SOURCE[0]}"
while [ -h "$SOURCE" ] ; do SOURCE="$(readlink "$SOURCE")"; done
DIR="$( cd -P "$( dirname "$SOURCE" )/.." && pwd )"

# Change into that directory
cd "$DIR" || exit

# Clean
echo "Cleaning.."
rm -f ./${BUILD_DIR}/**/*
mkdir -p ${BUILD_DIR}

if ! which gox > /dev/null; then
    echo "Installing gox..."
    go install github.com/mitchellh/gox@latest
fi

# Instruct gox to build statically linked binaries
export CGO_ENABLED=1

# Set module download mode to readonly to not implicitly update go.mod
export GOFLAGS="-mod=readonly"

# Download any module dependencies
go mod download

# Start build
echo "Building ${APP_NAME} ${VER_FLAG} for ${FLAG_OS} ${FLAG_ARCH}..."
gox \
    -output "${BUILD_DIR}/{{.OS}}_{{.Arch}}/${APP_NAME}-${VER_FLAG}" \
    -ldflags "${LD_FLAGS}" \
    -os="${FLAG_OS}" \
    -arch="${FLAG_ARCH}" \
    -osarch="${FLAG_EXCLUDE_OSARCH}" \
    .

echo "Build completed"
