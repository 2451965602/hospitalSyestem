// Code generated by hertz generator.

package epidemic

import (
	epidemic "Hospital/biz/model/epidemic"
	"Hospital/biz/pack"
	"Hospital/biz/service"
	"Hospital/pkg/errno"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// AddEpidemicCase .
// @router /epidemic/add [POST]
func AddEpidemicCase(ctx context.Context, c *app.RequestContext) {
	var err error
	var req epidemic.AddEpidemicCaseRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, errno.NewErrNo(errno.ParamMissingErrorCode, "param missing:"+err.Error()))
		return
	}

	resp := new(epidemic.AddEpidemicCaseResponse)
	err = service.NewEpidemicService(ctx, c).CreatEpidemicCase(&req)
	if err != nil {
		pack.SendFailResponse(c, errno.ConvertErr(err))
		return
	}
	resp.Base = pack.BuildBaseResp(errno.Success)
	pack.SendResponse(c, resp)
}

// QueryEpidemicCaseById .
// @router /epidemic/info [GET]
func QueryEpidemicCaseById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req epidemic.QueryEpidemicCaseByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, errno.NewErrNo(errno.ParamMissingErrorCode, "param missing:"+err.Error()))
		return
	}

	resp := new(epidemic.QueryEpidemicCaseByIdResponse)
	v, err := service.NewEpidemicService(ctx, c).QueryEpidemicCaseById(&req)
	if err != nil {
		pack.SendFailResponse(c, errno.ConvertErr(err))
		return
	}
	resp.Data = pack.EpidemicCase(v)
	resp.Base = pack.BuildBaseResp(errno.Success)
	pack.SendResponse(c, resp)
}

// QueryEpidemicCaseByPatient .
// @router /epidemic/info/patient [GET]
func QueryEpidemicCaseByPatient(ctx context.Context, c *app.RequestContext) {
	var err error
	var req epidemic.QueryEpidemicCaseByPatientRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		pack.SendFailResponse(c, errno.NewErrNo(errno.ParamMissingErrorCode, "param missing:"+err.Error()))
		return
	}
	resp := new(epidemic.QueryEpidemicCaseByPatientResponse)
	v, err := service.NewEpidemicService(ctx, c).QueryEpidemicCaseByPatient(&req)
	if err != nil {
		pack.SendFailResponse(c, errno.ConvertErr(err))
		return
	}

	resp.Data = pack.EpidemicCaseList(v)
	resp.Base = pack.BuildBaseResp(errno.Success)
	pack.SendResponse(c, resp)
}
