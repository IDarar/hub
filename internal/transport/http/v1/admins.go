package v1

import "github.com/gin-gonic/gin"

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

func (h *Handler) grantRole(c *gin.Context) {
	idFromCtx, _ := c.Get(userCtx)
	idStr, _ := idFromCtx.(string)

	c.String(200, string(idStr))
}
func (h *Handler) revokeRole(c *gin.Context) {
}
