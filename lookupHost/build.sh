#!/bin/bash

set -e

BUILD_DIR="build"
SRC="main.go"
OUT="clookup"
COMPILER="go"

build() {
    echo "building.."
    $COMPILER build -o $BUILD_DIR/$OUT $SRC 
    echo "completed. Find build in build dir"
}

build