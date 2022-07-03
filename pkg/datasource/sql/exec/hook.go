/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package exec

import (
	"context"

	"github.com/seata/seata-go-datasource/sql/types"
)

var (
	hookSolts = map[types.SQLType][]SQLInterceptor{}
)

// RegisHook not goroutine safe
func RegisHook(hook SQLInterceptor) {
	_, ok := hookSolts[hook.Type()]

	if !ok {
		hookSolts[hook.Type()] = make([]SQLInterceptor, 0, 4)
	}

	hookSolts[hook.Type()] = append(hookSolts[hook.Type()], hook)
}

// SQLHook SQL execution front and back interceptor
// case 1. Used to intercept SQL to achieve the generation of front and rear mirrors
// case 2. Burning point to report
// case 3. SQL black and white list
type SQLInterceptor interface {
	Type() types.SQLType

	// Before
	Before(ctx context.Context, txCtx *types.TransactionContext, query string, args ...interface{})

	// After
	After(ctx context.Context, txCtx *types.TransactionContext, query string, args ...interface{})
}