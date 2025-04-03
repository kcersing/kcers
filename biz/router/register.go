// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	admin_schedule "kcers/biz/router/admin/schedule"
	auth "kcers/biz/router/auth"
	banner "kcers/biz/router/banner"
	captcha "kcers/biz/router/captcha"
	contract "kcers/biz/router/contract"
	dictionary "kcers/biz/router/dictionary"
	entry "kcers/biz/router/entry"
	logs "kcers/biz/router/logs"
	member "kcers/biz/router/member"
	menu "kcers/biz/router/menu"
	order "kcers/biz/router/order"
	payment "kcers/biz/router/payment"
	product "kcers/biz/router/product"
	pub "kcers/biz/router/pub"
	service "kcers/biz/router/service"
	sms "kcers/biz/router/sms"
	sys "kcers/biz/router/sys"
	token "kcers/biz/router/token"
	user "kcers/biz/router/user"
	venue "kcers/biz/router/venue"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	venue.Register(r)

	user.Register(r)

	token.Register(r)

	sys.Register(r)

	sms.Register(r)

	service.Register(r)

	admin_schedule.Register(r)

	pub.Register(r)

	product.Register(r)

	payment.Register(r)

	order.Register(r)

	menu.Register(r)

	member.Register(r)

	logs.Register(r)

	entry.Register(r)

	dictionary.Register(r)

	contract.Register(r)

	captcha.Register(r)

	banner.Register(r)

	auth.Register(r)
}
