package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type DeleteManualBackupResponse struct {
	// 任务ID。

	JobId          *string `json:"job_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteManualBackupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteManualBackupResponse struct{}"
	}

	return strings.Join([]string{"DeleteManualBackupResponse", string(data)}, " ")
}
