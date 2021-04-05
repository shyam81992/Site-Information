package handler

import (
	"os"
	"testing"
)

func TestMain(t *testing.M) {
	exitVal := t.Run()
	os.Setenv("integration_testing", "true")

	if os.Getenv("integration_testing") == "true" {
		//config.LoadConfig()
	}
	os.Exit(exitVal)
}
