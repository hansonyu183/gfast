package router

import (
	//"gfast/hook"
	//"gfast/middleWare"
	//ap "gfast/app/api"
	"gfast/erp/api"
	"gfast/erp/desc"
	"gfast/erp/doc"
	"gfast/erp/list"
	"gfast/erp/query"
	"gfast/erp/vou"

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
		group.REST("/list/:name", list.New())
		group.REST("/vou/:id", vou.Vou)
		group.REST("/doc/:type/:id", doc.Doc)
		group.REST("/docAct/:type/:id/:act", doc.DocAct)
		group.REST("/desc/:user", desc.Desc)
		group.REST("/dbDesc/:table", desc.DBDesc)
		group.REST("/docList/:type", doc.DocList)
		group.REST("/docOption/:type", doc.DocOption)
		group.REST("/query/edtIo", query.EdtIo)
		group.REST("/query/edtIoSum", query.EdtIoSum)
		group.REST("/query/edtIoGroup", query.EdtIoGroup)
		group.REST("/query/edtIoGroupSum", query.EdtIoGroupSum)
	})

}
