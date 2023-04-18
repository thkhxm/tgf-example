package main

import "github.com/thkhxm/tgf/util"

// ***************************************************
// @Link  https://github.com/thkhxm/tgf
// @Link  https://gitee.com/timgame/tgf
// @QQ群 7400585
// author tim.huang<thkhxm@gmail.com>
// @Description
// 2023/4/18
// ***************************************************
func main() {
	util.SetExcelPath("../excel")
	util.SetExcelToGoPath("../common/conf")
	util.SetExcelToJsonPath("../common/conf/js")
	util.ExcelExport()
}
