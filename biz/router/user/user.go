// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	user "kcers/biz/handler/user"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_api := root.Group("/api", _apiMw()...)
		_api.POST("/register", append(_registerMw(), user.Register)...)
		{
			_admin := _api.Group("/admin", _adminMw()...)
			_admin.POST("/user", append(_deleteuserMw(), user.DeleteUser)...)
			_user := _admin.Group("/user", _userMw()...)
			_user.GET("/profile", append(_userprofileMw(), user.UserProfile)...)
			_user.POST("/set-default-venue", append(_setdefaultvenueMw(), user.SetDefaultVenue)...)
			_user.POST("/set-role", append(_setuserroleMw(), user.SetUserRole)...)
			_user.POST("/status", append(_updateuserstatusMw(), user.UpdateUserStatus)...)
			{
				_user0 := _admin.Group("/user", _user0Mw()...)
				_user0.POST("/change-password", append(_changepasswordMw(), user.ChangePassword)...)
				_user0.POST("/create", append(_createuserMw(), user.CreateUser)...)
				_user0.GET("/info", append(_userinfoMw(), user.UserInfo)...)
				_user0.POST("/list", append(_userlistMw(), user.UserList)...)
				_user0.POST("/perm", append(_userpermcodeMw(), user.UserPermCode)...)
				_user0.POST("/update", append(_updateuserMw(), user.UpdateUser)...)
			}
		}
	}
}
