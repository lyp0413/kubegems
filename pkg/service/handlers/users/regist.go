package userhandler

import (
	"github.com/gin-gonic/gin"
	"kubegems.io/pkg/service/handlers/base"
)

type UserHandler struct {
	base.BaseHandler
}

func (h *UserHandler) RegistRouter(rg *gin.RouterGroup) {
	rg.GET("/user", h.CheckIsSysADMIN, h.ListUser)
	rg.POST("/user", h.CheckIsSysADMIN, h.PostUser)
	rg.GET("/user/:user_id", h.CheckIsSysADMIN, h.RetrieveUser)
	rg.PUT("/user/:user_id", h.CheckIsSysADMIN, h.PutUser)
	rg.DELETE("/user/:user_id", h.CheckIsSysADMIN, h.DeleteUser)
	rg.GET("/user/:user_id/tenant", h.ListUserTenant)
	rg.POST("/user/:user_id/reset_password", h.CheckIsSysADMIN, h.ResetUserPassword)
	rg.GET("/user/_/environment/:environment_id", h.ListEnvironmentUser) // TODO: 严格来说，应该校验这些环境是否在用户当前的虚拟空间中
}
