/*
 * Copyright 2018 Marco Helmich
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/mhelmich/scripty/parser"
	"github.com/sirupsen/logrus"
	"llvm.org/git/llvm.git/bindings/go/llvm"
)

var shouldRun bool = true

func main() {
	logrus.Infof("llvm version: %s", llvm.Version)
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		cleanup()
		os.Exit(1)
	}()

	interpreterLoop()
}

func cleanup() {
	shouldRun = false
}

func interpreterLoop() {
	for shouldRun {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">>> ")
		text, _ := reader.ReadString('\n')
		parser.Parse(text)
	}
}
