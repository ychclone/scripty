#!/bin/sh

rm -f ../../parsergen/*
java -jar ./antlr-4.7-complete.jar -Dlanguage=Go scripty.g4 -o ../../parser/parsergen -package parsergen
