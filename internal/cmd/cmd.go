package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/os/gcmd"

	"demo/internal/controller/hello"
)

type CommonRes struct {
	g.Meta  `mime:"application/json"`
	Code    int         `json:"code" des:"Status Code"`
	Message string      `json:"message" des:"Status Message"`
	Data    interface{} `json:"data"`
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			openapi := s.GetOpenApi()
			openapi.Config.CommonResponse = CommonRes{}
			openapi.Config.CommonResponseDataField = `Data`
			openapi.Info = goai.Info{
				Title:       "CSSGEN API Reference",
				Description: "CSS generator",
			}
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
