package charttest_test

import (
	"os"
	"testing"
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
}
