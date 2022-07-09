/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-08 13:47:29
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-08 14:59:53
 * @FilePath: \ginx\models\alioss\uploadFileToOss.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

// TODO:这个文件可以不要
package alioss

import (
	"bytes"
	"ginx/config"
	"ginx/utils"
	"path"
)

func UploadToOss(fileName string, bucketName string, fileByte []byte) (url string, err error) {
	domain := config.Cfg.OssDomain
	bucket, err := GetBucket(bucketName)
	if err != nil {
		utils.HandleError(err)
	}

	folderName := utils.GetDay()
	yunFileTmpPath := path.Join("blog", folderName+"/"+fileName)

	err = bucket.PutObject(yunFileTmpPath, bytes.NewReader([]byte(fileByte)))
	if err != nil {
		utils.HandleError(err)
	}

	return domain + yunFileTmpPath, nil
}
