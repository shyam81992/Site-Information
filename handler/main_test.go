package handler

import (
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	exitVal := t.Run()
	if os.Getenv("INTEGRATION_TESTING") == "true" {
		//config.LoadConfig()
	}
	os.Exit(exitVal)
}
