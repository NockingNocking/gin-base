/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-07 23:04:38
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-07 23:06:51
 * @FilePath: \ginx\threadlocal\threadlocal.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package threadlocal

import "github.com/jtolds/gls"

var (
	Mgr = gls.NewContextManager()
	Rid = gls.GenSym()
)
