package Services

import (
	"gebi/app/Http/Serializer"
	"github.com/spf13/viper"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	ft "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ft/v20200304"
)

type FtChangeAgePicService struct {
	Age   int64  `json:"age" form:"age" binding:"required"`
	Image string `json:"image" form:"image" binding:"required"`
}

func (receiver FtChangeAgePicService) ChangeAgePic() string {
	client := ftclient()

	request := ft.NewChangeAgePicRequest()

	request.Image = common.StringPtr(receiver.Image)
	request.AgeInfos = []*ft.AgeInfo{
		{
			Age: common.Int64Ptr(receiver.Age),
		},
	}

	response, err := client.ChangeAgePic(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		Serializer.Err(50001, "An Ft API error has returned:", err)
	}

	return response.ToJsonString()
}

func ftclient() *ft.Client {
	secretId := viper.GetString("tencent_cloud.credential.secretId")
	secretKey := viper.GetString("tencent_cloud.credential.secretKey")

	credential := common.NewCredential(
		secretId,
		secretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "ft.tencentcloudapi.com"
	client, err := ft.NewClient(credential, "ap-guangzhou", cpf)
	if err != nil {
		Serializer.Err(50002, "An Ft Client new error:", err)
	}

	return client
}
