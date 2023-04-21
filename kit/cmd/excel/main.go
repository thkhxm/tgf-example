package main

import (
	"github.com/thkhxm/tgf"
	"github.com/thkhxm/tgf/component"
	"github.com/thkhxm/tgf/db"
	"github.com/thkhxm/tgf/example/common/conf"
	"github.com/thkhxm/tgf/util"
)

// ***************************************************
// @Link  https://github.com/thkhxm/tgf
// @Link  https://gitee.com/timgame/tgf
// @QQ群 7400585
// author tim.huang<thkhxm@gmail.com>
// @Description
// 2023/4/18
// ***************************************************
func main() {
	//
	db.WithCacheModule(tgf.CacheModuleClose)
	//
	util.SetExcelPath("./excel")
	util.SetExcelToGoPath("../common/conf")
	util.SetExcelToJsonPath("../common/conf/js")
	util.ExcelExport()

	//生成对应kv
	component.WithConfPath("../common/conf/js")
	//生成kv配置go文件
	component.InitGameConfToMem()
	codes := component.GetAllGameConf[*conf.ErrorCodeConf]()
	data := make([]util.TemplateKeyValueData, len(codes), len(codes))
	for i, code := range codes {
		data[i] = util.TemplateKeyValueData{
			FieldName: code.FieldName,
			Values:    code.Code,
		}
	}
	util.JsonToKeyValueGoFile("define", "error_code", "../common/define", "uint32", data)
}
