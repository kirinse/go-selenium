package integrationtests

import (
	"strings"
	"testing"
)

func Test_NavigateBack_NavigateBackWorksCorrectly(t *testing.T) {
	setUp()
	defer tearDown()

	driver := createDriver(t)
	_, err := driver.CreateSession()
	if err != nil {
		errorAndWrap(t, "Error thrown whilst creating session.", err)
	}

	goResp, err := driver.Go("https://news.ycombinator.com")
	if err != nil || goResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	goResp, err = driver.Go("https://google.co.uk")
	if err != nil || goResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating or result was not a success.", err)
	}

	backResp, err := driver.Back()
	if err != nil || backResp.State != "success" {
		errorAndWrap(t, "Error was thrown whilst navigating back or results was not a success.", err)
	}

	currentURLResp, err := driver.CurrentURL()
	if err != nil || !strings.HasPrefix(currentURLResp.URL, "https://news.ycombinator.com") {
		errorAndWrap(t, "Error was thrown or URL was not what it should have been.", err)
	}

	printObjectResult(currentURLResp)
}
