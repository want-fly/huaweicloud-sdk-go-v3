package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type CreateRecordRuleResponse struct {
	// 规则id，由服务端返回。创建或修改的时候不携带

	Id *string `json:"id,omitempty"`
	// 直播推流域名

	PublishDomain *string `json:"publish_domain,omitempty"`
	// 应用名，如果任意应用填写*。录制规则匹配的时候，优先精确app匹配，如果匹配不到，则匹配*

	App *string `json:"app,omitempty"`
	// 录制的流名，如果任意流名则填写*。录制规则匹配的时候，优先精确stream匹配，如果匹配不到，则匹配*

	Stream *string `json:"stream,omitempty"`
	// 录制类型，包括：CONTINUOUS_RECORD，COMMAND_RECORD，PLAN_RECORD, ON_DEMAND_RECORD。默认CONTINUOUS_RECORD。 - CONTINUOUS_RECORD: 持续录制，在该规则类型配置后，只要有流到推送到录制系统，就触发录制。 - COMMAND_RECORD: 命令录制，在该规则类型配置后，在流推送到录制系统后，租户需要通过命令控制该流的录制开始和结束。命令控制的接口参考/v1/{project_id}/record/control - PLAN_RECORD: 计划录制，在该规则类型配置后，推的流如果在计划录制的时间区间则触发录制。 - ON_DEMAND_RECORD: 按需录制，在该规则类型配置后，录制系统收到推流后，需要调用租户提供的接口查询录制规则，并根据规则录制。租户提供的接口定义参考：/customer-record-ondemand-url

	RecordType *CreateRecordRuleResponseRecordType `json:"record_type,omitempty"`

	DefaultRecordConfig *DefaultRecordConfig `json:"default_record_config,omitempty"`
	// 创建时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间。 在查询的时候返回

	CreateTime *string `json:"create_time,omitempty"`
	// 修改时间，格式：yyyy-mm-ddThh:mm:ssZ，UTC时间。 在查询的时候返回

	UpdateTime     *string `json:"update_time,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRecordRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRecordRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateRecordRuleResponse", string(data)}, " ")
}

type CreateRecordRuleResponseRecordType struct {
	value string
}

type CreateRecordRuleResponseRecordTypeEnum struct {
	CONTINUOUS_RECORD CreateRecordRuleResponseRecordType
	COMMAND_RECORD    CreateRecordRuleResponseRecordType
	PLAN_RECORD       CreateRecordRuleResponseRecordType
	ON_DEMAND_RECORD  CreateRecordRuleResponseRecordType
}

func GetCreateRecordRuleResponseRecordTypeEnum() CreateRecordRuleResponseRecordTypeEnum {
	return CreateRecordRuleResponseRecordTypeEnum{
		CONTINUOUS_RECORD: CreateRecordRuleResponseRecordType{
			value: "CONTINUOUS_RECORD",
		},
		COMMAND_RECORD: CreateRecordRuleResponseRecordType{
			value: "COMMAND_RECORD",
		},
		PLAN_RECORD: CreateRecordRuleResponseRecordType{
			value: "PLAN_RECORD",
		},
		ON_DEMAND_RECORD: CreateRecordRuleResponseRecordType{
			value: "ON_DEMAND_RECORD",
		},
	}
}

func (c CreateRecordRuleResponseRecordType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRecordRuleResponseRecordType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
