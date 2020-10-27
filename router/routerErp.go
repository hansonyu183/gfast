package router

import (
	//"gfast/hook"
	//"gfast/middleWare"
	"gfast/erp/api"
	"gfast/erp/dict"
	"gfast/erp/eba"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//后端路由处理
func init() {
	s := g.Server()
	group := s.Group("/")
	group.Group("/erp", func(group *ghttp.RouterGroup) {
		//group.Middleware(middleWare.Auth) //后台权限验证
		//后台操作日志记录
		//group.Hook("/*", ghttp.HOOK_AFTER_OUTPUT, hook.OperationLog)
		group.PATCH("/{.struct}/{.method}", api.NewWork())
		group.REST("/eba/list", eba.NewList())
		group.REST("/eba", eba.NewVou())
		group.REST("/dict", dict.NewDict())

	})
}
