package factory

import (
	"bet365/weChat/core/container"
	"bet365/weChat/global/my_errors"
	"bet365/weChat/global/variable"
	"bet365/weChat/http/validator/core/interf"
	"github.com/gin-gonic/gin"
)

// 表单参数验证器工厂（请勿修改）
func Create(key string) func(context *gin.Context) {

	if value := container.CreateContainersFactory().Get(key); value != nil {
		if val, isOk := value.(interf.ValidatorInterface); isOk {
			return val.CheckParams
		}
	}
	variable.ZapLog.Error(my_errors.ErrorsValidatorNotExists + ", 验证器模块：" + key)
	return nil
}
