package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListServiceTypesRequest struct {
	// 语言。zh_CN：中文en_US：英文缺省为zh_CN。

	XLanguage *string `json:"X-Language,omitempty"`
	// 云服务类型编码。例如OBS的云服务类型编码为“hws.service.type.obs”。

	ServiceTypeCode *string `json:"service_type_code,omitempty"`
}

func (o ListServiceTypesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServiceTypesRequest struct{}"
	}

	return strings.Join([]string{"ListServiceTypesRequest", string(data)}, " ")
}
