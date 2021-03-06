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

import "llvm.org/git/llvm.git/bindings/go/llvm"

func NewScopeContext(llvmCtx llvm.Context, module llvm.Module) *ScopeContext {
	return &ScopeContext{
		llvmCtx:   llvmCtx,
		module:    module,
		customCtx: make(map[string]interface{}),
	}
}

type ScopeContext struct {
	llvmCtx   llvm.Context
	module    llvm.Module
	customCtx map[string]interface{}
	parent    *ScopeContext
}

func (sc *ScopeContext) LlvmCtx() llvm.Context {
	return sc.llvmCtx
}

func (sc *ScopeContext) Module() llvm.Module {
	return sc.module
}

func (sc *ScopeContext) Set(key string, value interface{}) {
	sc.customCtx[key] = value
}

func (sc *ScopeContext) Get(key string) interface{} {
	csc := sc
	v, ok := sc.customCtx[key]
	for !ok && csc.parent != nil {
		csc = csc.parent
		v, ok = csc.customCtx[key]
	}
	return v
}

func (sc *ScopeContext) CreateChildScope() *ScopeContext {
	child := NewScopeContext(sc.llvmCtx, sc.module)
	child.parent = sc
	return child
}
