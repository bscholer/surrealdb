// Copyright © 2016 SurrealDB Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"context"

	"github.com/surrealdb/surrealdb/sql"
	"github.com/surrealdb/surrealdb/util/data"
)

func (e *executor) executeLet(ctx context.Context, stm *sql.LetStatement) (out []interface{}, err error) {

	var vars = ctx.Value(ctxKeyVars).(*data.Doc)

	switch what := stm.What.(type) {
	case *sql.Void:
		vars.Del(stm.Name.VA)
	case *sql.Empty:
		vars.Del(stm.Name.VA)
	default:
		val, err := e.fetch(ctx, what, nil)
		if err != nil {
			return nil, err
		}
		vars.Set(val, stm.Name.VA)
	}

	return

}
