package init

import (
	"fmt"
	"io/ioutil"

	"github.com/rerost/unstable/initialize/internal/config"
	api_pb "github.com/rerost/unstable/initialize/internal/proto"
)

const (
	ConfigFileName = "unstable.json"
)

var Cfg api_pb.Config

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

	Cfg = *cfg
}
