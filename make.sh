#!/bin/sh

LLVM_ROOT="$HOME/github/llvm-6.0.1.src/dist"
LLVM_CONFIG="${LLVM_ROOT}/bin/llvm-config"

export CGO_CPPFLAGS="$("$LLVM_CONFIG" --cppflags)"
export CGO_CXXFLAGS=-std=c++11
export CGO_LDFLAGS="$("$LLVM_CONFIG" --ldflags --libs --system-libs all)"

go build -tags byollvm -v
