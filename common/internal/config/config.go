package config

import (
	"encoding/json"

	api_pb "github.com/rerost/unstable/common/internal/proto"
	"github.com/srvc/fail"
)

func LoadConfig(configJSON []byte) (*api_pb.Config, error) {
	cfg := new(api_pb.Config)

	if err := json.Unmarshal(configJSON, cfg); err != nil {
		return nil, fail.Wrap(err)
	}

	return cfg, nil
}
