// Code generated by hertz generator. DO NOT EDIT.

package biobank

import (
	biobank "Hospital/biz/handler/biobank"
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
		_sample := root.Group("/sample", _sampleMw()...)
		_sample.POST("/add", append(_addsampleMw(), biobank.AddSample)...)
		_sample.GET("/info", append(_querysamplebyidMw(), biobank.QuerySampleById)...)
		_info := _sample.Group("/info", _infoMw()...)
		_info.GET("/patient", append(_querysamplebypatientMw(), biobank.QuerySampleByPatient)...)
	}
}
