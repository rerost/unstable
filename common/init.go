package common

import (
	"fmt"
	"io/ioutil"

	"github.com/rerost/unstable/common/internal/config"
	api_pb "github.com/rerost/unstable/common/internal/proto"
)

const (
	ConfigFileName = "unstable.json"
)

var Config api_pb.Config

func Init() {
	f, err := ioutil.ReadFile(ConfigFileName)
	if err != nil {
		panic(err)
	}

	cfg, err := config.LoadConfig(f)
	if err != nil {
		panic(err)
	}

	if cfg == nil {
		panic(fmt.Errorf("Failed to open config"))
	}

	Config = *cfg
}
