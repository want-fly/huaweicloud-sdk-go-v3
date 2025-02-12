package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 创建安全组规则请求体。
type CreateSecurityGroupRuleRequestBody struct {
	SecurityGroupRule *CreateSecurityGroupRuleOption `json:"security_group_rule"`
}

func (o CreateSecurityGroupRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateSecurityGroupRuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateSecurityGroupRuleRequestBody", string(data)}, " ")
}
