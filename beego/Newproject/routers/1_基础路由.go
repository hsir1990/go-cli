package routers

import (
	// "context"

	// "context"

	// "github.com/astaxie/beego/adapter/context"
	// "context"

	"github.com/beego/beego"
	cons "github.com/beego/beego/v2/context"
	// "github.com/beego/beego/v2/server/web/context"
)

func init() {
	beego.Get("/get", func(ctx *cons.C) {
		beego.Info("基础路由中的get请求")
		ctx.Output.Body([]byte("基础路由中的get请求，  get method"))
	})
}
