package router

import (
	//"gfast/hook"
	//"gfast/middleWare"
	//ap "gfast/app/api"
	"gfast/erp/api"
	"gfast/erp/doc"
	"gfast/erp/query"
	"gfast/erp/ui"
	"gfast/erp/vr"

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

		group.REST("/ui/desc/:user", ui.Desc)
		group.REST("/ui/opt/*name", ui.Opt)
		group.REST("/ui/auth/:user", ui.Auth)
		group.REST("/ui/dict/", ui.Dict)

		group.REST("/vr/:type/:id", vr.Vr)
		group.REST("/vrList/:type", vr.VrList)

		group.REST("/doc/:type/:id", doc.Doc)
		group.REST("/docList/:type", doc.DocList)
		group.REST("/docSub/:docType/:docID", doc.DocSub)
		group.REST("/query/edtIo", query.EdtIo)
		group.REST("/query/edtIoSum", query.EdtIoSum)
		group.REST("/query/edtIoGroup", query.EdtIoGroup)
		group.REST("/query/edtIoGroupSum", query.EdtIoGroupSum)
	})

}
