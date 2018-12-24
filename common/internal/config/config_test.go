package config_test

import (
	"testing"

	"github.com/rerost/unstable/common/internal/config"
)

func TestLoadConfig(t *testing.T) {
	t.Parallel()
	inputOutputPairs := []struct {
		name  string
		input string
	}{
		{
			name: "full filled",
			input: `
{
	"interval": 1,
	"server_error": true,
	"slow_response_option": {
		"enable": true,
		"time": 3
	},
	"server_error_option": {
		"enable": true
	}
}`,
		},
		{
			name: "empty",
			input: `
{
}`,
		},
		{
			name: "with invalid_field",
			input: `
{
	"invalid_field": true
}`,
		},
	}

	for _, p := range inputOutputPairs {
		t.Run(p.name, func(t *testing.T) {
			_, err := config.LoadConfig([]byte(p.input))
			if err != nil {
				t.Fail()
			}
		})
	}
}
