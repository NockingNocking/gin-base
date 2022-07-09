/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-07 12:22:23
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-08 14:40:09
 * @FilePath: \ginx\utils\utils.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package utils

import (
	"time"
)

// 获取当前时间戳
func GetUnix() int64 {
	return time.Now().Unix()
}

// 时间戳转换成时间
func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-01 00:00:00")
}

// 时间转换成时间戳
func DateToUnix(str string) int64 {
	template := "2006-01-02 15:00:00"
	t, err := time.ParseInLocation(template, str, time.Local)
	if err != nil {
		return 0
	}
	return t.Unix()
}

// 获取年月日
func GetDay() string {
	template := "20060101"
	return time.Now().Format(template)
}

// 获取当前日期
func GetDate() string {
	template := "2006-01-02 15:00:00"
	return time.Now().Format(template)
}

//处理错误
func HandleError(err error) error {
	if err == nil {
		return nil
	}
	return err
}

// 检查图片格式是否正确
func CheckFileExt(extName string) (string, bool) {
	allowExtMap := map[string]bool{
		".jpg":  true,
		".png":  true,
		".gif":  true,
		".jpeg": true,
	}

	if _, ok := allowExtMap[extName]; !ok {
		return "图片格式不对!", false
	}
	return "", true
}
