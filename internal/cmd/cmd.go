package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"

	"gf_demo/internal/controller/hello"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			// 使用Group方法创建分组路由,
			s.Group("/", func(group *ghttp.RouterGroup) {
				// 通过Middleware方法注册了一个中间件，该中间件是HTTP Server组件用于规范化路由的数据返回。
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				// 通过Bind方法的规范化路由方式绑定一个Hello路由对象，该路由对象下的所有公开方法均会被自动注册会路由。
				group.Bind(
					hello.NewV1(),
				)
			})
			s.Run()
			return nil
		},
	}
)
