// +build integration

package integration_test

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStatus(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	res, err := http.Get(address + "/version")
	require.NoError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	require.NoError(err)
	assert.Equal(
		`version="integration-tests-version" buildDate="integration-tests-build-date" serviceName="integration-tests-service-name" commit="integration-tests-commit-hash"`,
		string(body))
}
