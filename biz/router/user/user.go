// Code generated by hertz generator. DO NOT EDIT.

package user

import (
	user "Hospital/biz/handler/user"
	"github.com/cloudwego/hertz/pkg/app/server"
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
		_user := root.Group("/user", _userMw()...)
		_user.POST("/add", append(_newuserMw(), user.NewUser)...)
		_user.GET("/info", append(_queryMw(), user.Query)...)
		_info := _user.Group("/info", _infoMw()...)
		_info.GET("/list", append(_queryuserlistMw(), user.QueryUserList)...)
		_user.POST("/login", append(_loginMw(), user.Login)...)
		{
			_admin := _user.Group("/admin", _adminMw()...)
			_admin.PUT("/update", append(_updateadminMw(), user.UpdateAdmin)...)
		}
		{
			_doctor := _user.Group("/doctor", _doctorMw()...)
			_doctor.PUT("/update", append(_updatedoctorMw(), user.UpdateDoctor)...)
		}
		{
			_nurse := _user.Group("/nurse", _nurseMw()...)
			_nurse.PUT("/update", append(_updatenurseMw(), user.UpdateNurse)...)
		}
	}
}
