// Licensed to Elasticsearch B.V under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.
//
// Code generated, DO NOT EDIT

package elasticsearch_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
)

var (
	_ = fmt.Printf
	_ = os.Stdout
	_ = elasticsearch.NewDefaultClient
)

// <https://github.com/elastic/elasticsearch/blob/master/docs/reference/docs/reindex.asciidoc#L224>
//
// --------------------------------------------------------------------------------
// GET _refresh
// POST new_twitter/_search?size=0&filter_path=hits.total
// --------------------------------------------------------------------------------

func Test_docs_reindex_3ae03ba3b56e5e287953094050766738(t *testing.T) {
	es, _ := elasticsearch.NewDefaultClient()

	// tag:3ae03ba3b56e5e287953094050766738[]
	{
		res, err := es.Indices.Refresh()
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}

	{
		res, err := es.Search(
			es.Search.WithIndex("new_twitter"),
			es.Search.WithFilterPath("hits.total"),
			es.Search.WithSize(0),
			es.Search.WithPretty(),
		)
		fmt.Println(res, err)
		if err != nil { // SKIP
			t.Fatalf("Error getting the response: %s", err) // SKIP
		} // SKIP
		defer res.Body.Close() // SKIP
	}
	// end:3ae03ba3b56e5e287953094050766738[]
}
