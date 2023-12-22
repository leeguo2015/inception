package utils

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
)

type TencentOSSUploader struct {
	AppID      string
	SecretID   string
	SecretKey  string
	Region     string
	BucketName string
}

// 上传文件到腾讯云对象存储（COS）
func UploadFileToCos(BucketURL string, f io.Reader, secretID string, secretKey string) error {
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse(BucketURL)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 COS_SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: secretID,
			// 环境变量 COS_SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: secretKey,
			// Debug 模式，把对应 请求头部、请求内容、响应头部、响应内容 输出到标准输出
			Transport: &debug.DebugRequestTransport{
				RequestHeader: true,
				// Notice when put a large file and set need the request body, might happend out of memory error.
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})

	// Case1 上传对象
	name := time.Now().Format("2006/01/02")
	if _, err := c.Object.Put(context.Background(), name, f, nil); err != nil {
		return err
	}
	return nil
	// Case2 使用options上传对象
	// f = strings.NewReader("test xxx")
	// opt := &cos.ObjectPutOptions{
	// 	ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
	// 		ContentType: "text/html",
	// 	},
	// 	ACLHeaderOptions: &cos.ACLHeaderOptions{
	// 		//XCosACL: "public-read",
	// 		XCosACL: "private",
	// 	},
	// }
	// _, err = c.Object.Put(context.Background(), name, f, opt)
	// log_status(err)

	// // Case3 通过本地文件上传对象
	// _, err = c.Object.PutFromFile(context.Background(), name, "./test", nil)
	// log_status(err)

	// // Case4 查看上传进度
	// opt.ObjectPutHeaderOptions.Listener = &cos.DefaultProgressListener{}
	// _, err = c.Object.PutFromFile(context.Background(), name, "./test", opt)
}