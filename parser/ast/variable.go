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

package ast

import (
	"github.com/sirupsen/logrus"
	"llvm.org/git/llvm.git/bindings/go/llvm"
)

type Variable struct {
	Name string
}

func (v *Variable) GenCode(sc *ScopeContext, builder llvm.Builder) llvm.Value {
	val := sc.Get(v.Name)
	if val == nil {
		logrus.Errorf("Can't find variable '%s'", v.Name)
		return llvm.Value{}
	}

	return val.(llvm.Value)
}
