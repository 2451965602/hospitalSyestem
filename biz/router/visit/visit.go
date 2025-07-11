// Code generated by hertz generator. DO NOT EDIT.

package visit

import (
	visit "Hospital/biz/handler/visit"
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
		_visit := root.Group("/visit", _visitMw()...)
		_visit.POST("/add", append(_addvisitmessageMw(), visit.AddVisitMessage)...)
		_visit.GET("/info", append(_queryvisitmessageMw(), visit.QueryVisitMessage)...)
	}
}
