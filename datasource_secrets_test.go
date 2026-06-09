package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
)

func TestDataSourceSecretSchema(t *testing.T) {
	s := dataSourceSecret().Schema

	def, ok := s["default"]
	if !ok {
		t.Fatal("expected a \"default\" attribute on the secret data source")
	}

	if def.Type != schema.TypeString {
		t.Errorf("expected \"default\" to be a string, got: %s", def.Type)
	}

	if !def.Optional {
		t.Error("expected \"default\" to be optional")
	}

	if def.Required {
		t.Error("expected \"default\" to not be required")
	}
}
