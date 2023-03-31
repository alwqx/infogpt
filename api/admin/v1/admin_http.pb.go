// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.12
// source: admin/v1/admin.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationAdminAppInfo = "/admin.v1.Admin/AppInfo"
const OperationAdminBookSummary = "/admin.v1.Admin/BookSummary"
const OperationAdminHealthCheck = "/admin.v1.Admin/HealthCheck"
const OperationAdminOpenaiChat = "/admin.v1.Admin/OpenaiChat"
const OperationAdminUrlSummary = "/admin.v1.Admin/UrlSummary"

type AdminHTTPServer interface {
	// AppInfo Sends appinfo
	AppInfo(context.Context, *AppInfoRequest) (*AppInfoReply, error)
	// BookSummary book summary using openai
	BookSummary(context.Context, *SummaryReuqest) (*SummaryReply, error)
	// HealthCheck Sends a greeting
	HealthCheck(context.Context, *HealthRequest) (*HealthReply, error)
	// OpenaiChat proxy chat to openai
	OpenaiChat(context.Context, *OpenaiChatReuqest) (*OpenaiChatReply, error)
	// UrlSummary url summary using openai
	UrlSummary(context.Context, *SummaryReuqest) (*SummaryReply, error)
}

func RegisterAdminHTTPServer(s *http.Server, srv AdminHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/health", _Admin_HealthCheck0_HTTP_Handler(srv))
	r.GET("/v1/appinfo", _Admin_AppInfo0_HTTP_Handler(srv))
	r.POST("/v1/chat", _Admin_OpenaiChat0_HTTP_Handler(srv))
	r.POST("/v1/summary/url", _Admin_UrlSummary0_HTTP_Handler(srv))
	r.POST("/v1/summary/book", _Admin_BookSummary0_HTTP_Handler(srv))
}

func _Admin_HealthCheck0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in HealthRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminHealthCheck)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.HealthCheck(ctx, req.(*HealthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*HealthReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_AppInfo0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AppInfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminAppInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AppInfo(ctx, req.(*AppInfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AppInfoReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_OpenaiChat0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in OpenaiChatReuqest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminOpenaiChat)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.OpenaiChat(ctx, req.(*OpenaiChatReuqest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*OpenaiChatReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_UrlSummary0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SummaryReuqest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminUrlSummary)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UrlSummary(ctx, req.(*SummaryReuqest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SummaryReply)
		return ctx.Result(200, reply)
	}
}

func _Admin_BookSummary0_HTTP_Handler(srv AdminHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SummaryReuqest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAdminBookSummary)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.BookSummary(ctx, req.(*SummaryReuqest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SummaryReply)
		return ctx.Result(200, reply)
	}
}

type AdminHTTPClient interface {
	AppInfo(ctx context.Context, req *AppInfoRequest, opts ...http.CallOption) (rsp *AppInfoReply, err error)
	BookSummary(ctx context.Context, req *SummaryReuqest, opts ...http.CallOption) (rsp *SummaryReply, err error)
	HealthCheck(ctx context.Context, req *HealthRequest, opts ...http.CallOption) (rsp *HealthReply, err error)
	OpenaiChat(ctx context.Context, req *OpenaiChatReuqest, opts ...http.CallOption) (rsp *OpenaiChatReply, err error)
	UrlSummary(ctx context.Context, req *SummaryReuqest, opts ...http.CallOption) (rsp *SummaryReply, err error)
}

type AdminHTTPClientImpl struct {
	cc *http.Client
}

func NewAdminHTTPClient(client *http.Client) AdminHTTPClient {
	return &AdminHTTPClientImpl{client}
}

func (c *AdminHTTPClientImpl) AppInfo(ctx context.Context, in *AppInfoRequest, opts ...http.CallOption) (*AppInfoReply, error) {
	var out AppInfoReply
	pattern := "/v1/appinfo"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminAppInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) BookSummary(ctx context.Context, in *SummaryReuqest, opts ...http.CallOption) (*SummaryReply, error) {
	var out SummaryReply
	pattern := "/v1/summary/book"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminBookSummary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) HealthCheck(ctx context.Context, in *HealthRequest, opts ...http.CallOption) (*HealthReply, error) {
	var out HealthReply
	pattern := "/v1/health"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAdminHealthCheck))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) OpenaiChat(ctx context.Context, in *OpenaiChatReuqest, opts ...http.CallOption) (*OpenaiChatReply, error) {
	var out OpenaiChatReply
	pattern := "/v1/chat"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminOpenaiChat))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AdminHTTPClientImpl) UrlSummary(ctx context.Context, in *SummaryReuqest, opts ...http.CallOption) (*SummaryReply, error) {
	var out SummaryReply
	pattern := "/v1/summary/url"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAdminUrlSummary))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}