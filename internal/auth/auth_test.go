package auth

import (
	"testing"
)

func TestGetApiKey(t *testing.T) {
	testHeader := make(map[string][]string)

	testHeader["Authorization"] = []string{"ApiKey TestApiKey"}

	Want := "TestApiKey"
	got, err := GetAPIKey(testHeader)

	if err != nil {
		t.Errorf("Expected: %v, got: %v", Want, got)
	}
	if Want != got {
		t.Fatalf("expected: %v, got: %v", Want, got)
	}

}
