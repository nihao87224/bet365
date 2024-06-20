package cur_userinfo

import (
	"bet365/weChat/global/variable"
	"bet365/weChat/http/middleware/my_jwt"
	"github.com/gin-gonic/gin"
)

// GetCurrentUserId 获取当前用户的id
// @context 请求上下文
func GetCurrentUserId(context *gin.Context) (int64, bool) {
	tokenKey := variable.ConfigYml.GetString("Token.BindContextKeyName")
	currentUser, exist := context.MustGet(tokenKey).(my_jwt.CustomClaims)
	return currentUser.UserId, exist
}
