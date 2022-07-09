/*
 * @Author: Mocking 497773732@qq.com
 * @Date: 2022-07-08 12:34:26
 * @LastEditors: Mocking 497773732@qq.com
 * @LastEditTime: 2022-07-08 13:45:26
 * @FilePath: \ginx\models\alioss\ossInit.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package alioss

import (
	"ginx/config"
	"ginx/utils"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// 创建bucket
func CreateBucket(bucketName string) (string, error) {
	client, err := oss.New(config.Cfg.OssEndpoint, config.Cfg.OssAccessID, config.Cfg.OssAccessKey)
	if err != nil {
		utils.HandleError(err)
	}

	err = client.CreateBucket(bucketName)
	if err != nil {
		utils.HandleError(err)
	}
	return bucketName + "创建成功！", nil
}

// 删除bucket
func DeleteBucket(bucketName string) (string, error) {
	client, err := oss.New(config.Cfg.OssEndpoint, config.Cfg.OssAccessID, config.Cfg.OssAccessKey)
	if err != nil {
		utils.HandleError(err)
	}
	err = client.DeleteBucket(bucketName)
	if err != nil {
		utils.HandleError(err)
	}
	return bucketName + "删除成功！", nil
}

// 获取bucket
func GetBucket(bucketName string) (*oss.Bucket, error) {

	client, err := oss.New(config.Cfg.OssEndpoint, config.Cfg.OssAccessID, config.Cfg.OssAccessKey)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
