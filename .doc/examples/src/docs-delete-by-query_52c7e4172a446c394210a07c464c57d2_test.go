// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/frikky/go-elasticsearch/v8"
	"github.com/frikky/go-elasticsearch/v8/esapi"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/delete-by-query.asciidoc#L579>
//
// --------------------------------------------------------------------------------
// POST _delete_by_query/r1A2WoRbTwKZ516z6NEs5A:36619/_rethrottle?requests_per_second=-1
// --------------------------------------------------------------------------------

func Test_docs_delete_by_query_52c7e4172a446c394210a07c464c57d2(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:52c7e4172a446c394210a07c464c57d2[]
	res, err := es.DeleteByQueryRethrottle(
		"r1A2WoRbTwKZ516z6NEs5A:36619",
		esapi.IntPtr(-1),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:52c7e4172a446c394210a07c464c57d2[]
}
