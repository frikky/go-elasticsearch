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
	"strings"
	"testing"

	"github.com/frikky/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L784>
//
// --------------------------------------------------------------------------------
// POST _reindex
// {
//   "source": {
//     "index": "metricbeat-*"
//   },
//   "dest": {
//     "index": "metricbeat"
//   },
//   "script": {
//     "lang": "painless",
//     "source": "ctx._index = 'metricbeat-' + (ctx._index.substring('metricbeat-'.length(), ctx._index.length())) + '-1'"
//   }
// }
// --------------------------------------------------------------------------------

func Test_docs_reindex_973a3ff47fc4ce036ecd9bd363fef9f7(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:973a3ff47fc4ce036ecd9bd363fef9f7[]
	res, err := es.Reindex(
		strings.NewReader(`{
		  "source": {
		    "index": "metricbeat-*"
		  },
		  "dest": {
		    "index": "metricbeat"
		  },
		  "script": {
		    "lang": "painless",
		    "source": "ctx._index = 'metricbeat-' + (ctx._index.substring('metricbeat-'.length(), ctx._index.length())) + '-1'"
		  }
		}`))
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:973a3ff47fc4ce036ecd9bd363fef9f7[]
}
