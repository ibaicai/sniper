package server

import (
	"net/http"

	"sniper/cmd/server/hook"

	"github.com/bilibili/twirp"

	user_v1 "sniper/rpc/user/v1"
	"sniper/server/userserver1"
)

var hooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewLog(),
)

var loginHooks = twirp.ChainHooks(
	hook.NewRequestID(),
	hook.NewCheckLogin(),
	hook.NewLog(),
)

func initMux(mux *http.ServeMux, isInternal bool) {
	{
		server := &userserver1.Server{}
		handler := user_v1.NewUserServer(server, hooks)
		mux.Handle(user_v1.UserPathPrefix, handler)
	}
}

func initInternalMux(mux *http.ServeMux) {
}
