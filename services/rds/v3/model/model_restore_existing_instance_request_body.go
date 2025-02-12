package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RestoreExistingInstanceRequestBody struct {
	Source *RestoreToExistingInstanceRequestBodySource `json:"source"`

	Target *RestoreToExistingInstanceRequestBodyTarget `json:"target"`
}

func (o RestoreExistingInstanceRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestoreExistingInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"RestoreExistingInstanceRequestBody", string(data)}, " ")
}
