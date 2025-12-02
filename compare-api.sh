#!/bin/sh

set -e

env GOOS=linux GOARCH=arm64  go doc ./wgpu > /tmp/native.txt
env GOOS=js    GOARCH=wasm   go doc ./wgpu > /tmp/js.txt

diff /tmp/native.txt /tmp/js.txt
