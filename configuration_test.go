package main

import (
	"os"
	"testing"
)

func TestExtractConfiguration(t *testing.T) {
	// Set environment variables required for the test
	os.Setenv("DB_TYPE", "postgres")
	os.Setenv("DB_CONNECTION", "postgres://user:password@localhost:5432/dbname")
	os.Setenv("SERVER_ADDRESSS", "8080")
	os.Setenv("PRODUCTION", "false")

	// Call the function being tested
	cfg, err := ExtractConfiguration()
	if err != nil {
		t.Fatalf("Unexpected error while extracting configuration: %v", err)
	}

	// Assert that the expected values were loaded from the environment variables
	expectedCfg := ServiceConfig{
		Databasetype:  DBTYPE("postgres"),
		DBConnection:  "postgres://user:password@localhost:5432/dbname",
		ServerAddress: "8080",
		IsProduction:  "false",
	}
	if cfg != expectedCfg {
		t.Fatalf("Unexpected configuration. Expected %v, got %v", expectedCfg, cfg)
	}
}
