// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	admin_sys "kcers/biz/router/admin/sys"
	logs "kcers/biz/router/logs"
	token "kcers/biz/router/token"
	user "kcers/biz/router/user"
	venue "kcers/biz/router/venue"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	admin_sys.Register(r)

	logs.Register(r)

	venue.Register(r)

	user.Register(r)

	token.Register(r)
}
