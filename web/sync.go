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

package web

import (
	"github.com/surrealdb/fibre"
	"github.com/surrealdb/surrealdb/cnf"
	"github.com/surrealdb/surrealdb/db"
)

func syncer(c *fibre.Context, sync bool) (err error) {

	if c.Get("auth").(*cnf.Auth).Kind != cnf.AuthKV {
		return fibre.NewHTTPError(401)
	}

	c.Response().Header().Set("Content-Type", "application/octet-stream")

	switch sync {
	case true:
		return db.Sync(c.Response())
	case false:
		return db.Sync(c.Request().Body)
	}

	return nil

}
