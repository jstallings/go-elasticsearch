// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/update.asciidoc#L153>
//
// --------------------------------------------------------------------------------
// POST test/_update/1
// {
//     "script" : "ctx._source.new_field = 'value_of_new_field'"
// }
// --------------------------------------------------------------------------------

func Test_docs_update_eb30ba547e4a7b8f54f33ab259aca523(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:eb30ba547e4a7b8f54f33ab259aca523[]
	res, err := es.Update(
		"test",
		"1",
		strings.NewReader(`{
		  "script": "ctx._source.new_field = 'value_of_new_field'"
		}`),
		es.Update.WithPretty(),
	)
	fmt.Println(res, err)
	if err != nil { // SKIP
		t.Fatalf("Error getting the response: %s", err) // SKIP
	} // SKIP
	defer res.Body.Close() // SKIP
	// end:eb30ba547e4a7b8f54f33ab259aca523[]
}
