package charttest_test

import (
	"os"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func TestEnvs(t *testing.T) {
	sec := os.Getenv("TEST_SECRET")
	en := os.Getenv("TEST_ENV")
	if en == "" {
		t.Errorf("expected env to be set, but was: '%v'", en)
	}
	if sec == "" {
		t.Errorf("expected sec to be set, but was: '%v'", sec)
	}
	assert.NotEmpty(t, sec)
	assert.NotEmpty(t, en)
}
