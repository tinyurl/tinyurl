package server

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortenURL(t *testing.T) {
	startTestServer(t)
	shortenURLAPI := TestAddr + "/api/v1/shorten"

	// test origin not existed in db
	postData1 := url.Values{}
	postData1.Set("origin_url", TestOriginURL)
	ret1 := PostForm(shortenURLAPI, postData1)
	map1 := ret1.(map[string]interface{})
	assert.NotNil(t, map1)
	shortPath := map1["short_path"].(string)
	assert.Equal(t, map1["message"], ShortenURLSuccess)

	// test url has existed
	testURL := newTestURL()
	testURL.ShortPath = shortPath
	updateTestURL(testURL)
	ret2 := PostForm(shortenURLAPI, postData1)
	map2 := ret2.(map[string]interface{})
	assert.NotNil(t, map2)
	assert.Equal(t, map2["message"], ShortPathExisted)
	assert.Equal(t, map2["short_path"], shortPath)

	clearDatabase()
}
