package v1

import (
	"net/http"

	"github.com/IDarar/hub/pkg/logger"
	"github.com/gin-gonic/gin"
)

//35.10
func (h *Handler) initAdminsRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins", h.adminIdentity)
	{
		roles := admins.Group("/roles")
		{
			roles.POST("/grant-role", h.grantRole)
			roles.POST("/revoke-role", h.revokeRole)
		}

	}

}

type grantRoleInput struct {
	UserName string `json:"username"  binding:"required,min=2,max=64"`
	Role     string `json:"role"  binding:"required,min=5,max=64"`
}

// @Summary	admin GrantRole
// @Security AdminAuth
// @Tags roles
// @Description admin-grantrole
// @ModuleID admin
// @Accept  json
// @Produce  json
// @Param input body grantRoleInput true "role granting info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/roles/grant-role [post]
func (h *Handler) grantRole(c *gin.Context) {
	roles, ex := c.Get(roleCtx)
	if !ex {
		newResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}
	var inp grantRoleInput

	if err := c.BindJSON(&inp); err != nil {
		logger.Debug("bug")
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Admin.GrantRole(inp.UserName, inp.Role, roles)
	if err != nil {
		logger.Debug("bug")

		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) revokeRole(c *gin.Context) {
}
