#!/usr/bin/bash

# To install bazelisk with "go", simply;
# go install github.com/bazelbuild/bazelisk@latest

# To reset, run "--migrate"



source build/envsetup.sh
export PATH=~/.bin:$PATH
export GOPATH=~/GO
export ANDROID_BUILD_TOP=/root/WORKSPACE
export SDCLANG="true"
bash build/make/rbesetup.sh
