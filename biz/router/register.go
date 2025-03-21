// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	auth "kcers/biz/router/auth"
	logs "kcers/biz/router/logs"
	menu "kcers/biz/router/menu"
	sys "kcers/biz/router/sys"
	token "kcers/biz/router/token"
	user "kcers/biz/router/user"
	venue "kcers/biz/router/venue"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	menu.Register(r)

	auth.Register(r)

	logs.Register(r)

	token.Register(r)

	user.Register(r)

	venue.Register(r)

	sys.Register(r)
}
