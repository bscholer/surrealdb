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

package sql

func (p *parser) parseReturnStatement() (stmt *ReturnStatement, err error) {

	grw := p.buf.rw

	p.buf.rw = false

	defer func() {
		p.buf.rw = grw
	}()

	stmt = &ReturnStatement{}

	// The next query part can be any expression
	// including a parenthesised expression or a
	// binary expression so handle accordingly.

	stmt.What, err = p.parseWhat()
	if err != nil {
		return nil, err
	}

	// If this query has any subqueries which
	// need to alter the database then mark
	// this query as a writeable statement.

	stmt.RW = p.buf.rw

	return

}
