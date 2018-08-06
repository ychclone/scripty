#!/bin/sh

# LLVM_ROOT="$(brew --prefix llvm)"
# LLVM_ROOT="/usr/local/opt/llvm"
# LLVM_ROOT="/Users/marco.helmich/github/llvm/dist"
LLVM_ROOT="/Users/marco.helmich/github/llvm-6.0.1.src/dist"
LLVM_CONFIG="${LLVM_ROOT}/bin/llvm-config"

export CGO_CPPFLAGS="$("$LLVM_CONFIG" --cppflags)"
export CGO_CXXFLAGS=-std=c++11
export CGO_LDFLAGS="$("$LLVM_CONFIG" --ldflags --libs --system-libs all)"
# export GO_EXTLINK_ENABLED=1

go build -tags byollvm -v
