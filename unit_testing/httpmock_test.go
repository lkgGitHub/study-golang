package unit_testing

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestFetchArticles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles",
		httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Article"}]`))

	// Regexp match (could use httpmock.RegisterRegexpResponder instead)
	httpmock.RegisterResponder("GET", `=~^https://api\.mybiz\.com/articles/id/\d+\z`,
		httpmock.NewStringResponder(200, `{"id": 1, "name": "My Great Article"}`))

	// do stuff that makes a request to articles

	// get count info
	httpmock.GetTotalCallCount()

	// get the amount of calls for the registered responder
	info := httpmock.GetCallCountInfo()
	if info["GET https://api.mybiz.com/articles"] == 1 { // number of GET calls made to https://api.mybiz.com/articles

	}
	if info["GET https://api.mybiz.com/articles/id/12"] == 1 { // number of GET calls made to https://api.mybiz.com/articles/id/12
	}
	if info[`GET =~^https://api\.mybiz\.com/articles/id/\d+\z`] == 1 { // number of GET calls made to https://api.mybiz.com/articles/id/<any-number>
	}
}
