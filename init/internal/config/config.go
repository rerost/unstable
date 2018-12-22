package config

import (
	"encoding/json"

	api_pb "github.com/rerost/unstable/init/internal/proto"
	"github.com/srvc/fail"
)

const (
	FileName = "unstable.json"
)

func LoadConfig(configJSON []byte) (*api_pb.Config, error) {
	cfg := new(api_pb.Config)

	if err := json.Unmarshal(configJSON, cfg); err != nil {
		return nil, fail.Wrap(err)
	}

	return cfg, nil
}
