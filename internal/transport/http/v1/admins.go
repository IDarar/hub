package v1

import "github.com/gin-gonic/gin"

//35.10
func (h *Handler) initAdminsRoutes(api *gin.RouterGroup) {
	admins := api.Group("/admins") //51.23
	{
		admins.POST("/grant-role", h.grantRole)

		admins.POST("/revoke-role", h.revokeRole)

	}
}
func (h *Handler) grantRole(c *gin.Context) {

}
func (h *Handler) revokeRole(c *gin.Context) {
}
