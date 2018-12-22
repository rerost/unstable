package init

var Cfg config

func Init() {
	Cfg := LoadConfig()
}
