package uhttp

import (
	"net/http"
	"time"

	"github.com/rerost/unstable/common"
)

func init() {
	common.Init()
}

func WithUnstable(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cfg := common.Config
		if cfg.Interval != 0 && uint64(time.Now().Second())%cfg.Interval == 0 {
			if cfg.SlowResponseOption.GetEnable() {
				time.Sleep(time.Duration(cfg.SlowResponseOption.GetTime() * 1000000))
			}
			if cfg.ServerErrorOption.GetEnable() {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Fail by unstable"))
				return
			}
		}
		next.ServeHTTP(w, r)
	}
}
