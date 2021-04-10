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
		content := admins.Group("/content")
		{

			content.POST("", h.createTreatise)
			content.PUT("/:id", h.updateTreatise)
			content.DELETE("/:id", h.deleteTreatise)
			content.POST("/:id/parts", h.createPart)
			//For treatises without parts division
			content.POST("/:id/proposition", h.createProposition)

		}
		/*parts := content.Group("/parts")
		{
			parts.POST("/:id")
			parts.PUT("/:id")
			parts.PUT("/:id")
			//For treatises divided into parts
			parts.POST("/:id/proposition")
		}
		propositions := content.Group("/propositions")
		{
			propositions.POST("/:id")
			propositions.PUT("/:id")
			propositions.PUT("/:id")
		}*/

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
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Admin.GrantRole(inp.UserName, inp.Role, roles)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

//TODO
type revokeRoleInput struct {
	UserName string `json:"username"  binding:"required,min=2,max=64"`
	Role     string `json:"role"  binding:"required,min=5,max=64"`
}

func (h *Handler) revokeRole(c *gin.Context) {
}

type treatiseCreateInput struct {
	ID          string `json:"id"  binding:"required"`
	Title       string `json:"title"  binding:"required"`
	Description string `json:"description"  binding:"required"`
	Date        string `json:"date"  binding:"required"`
}

// @Summary	admin CreateTreatise
// @Security AdminAuth
// @Tags content
// @Description CreateTreatise
// @ModuleID createTreatise
// @Accept  json
// @Produce  json
// @Param input body treatiseCreateInput true "treatise info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/content [post]
func (h *Handler) createTreatise(c *gin.Context) {
	var inp treatiseCreateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	roles, ex := c.Get(roleCtx)
	if !ex {
		newResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err := h.services.Content.Create(inp.ID, inp.Title, inp.Date, inp.Description, roles)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}
func (h *Handler) updateTreatise(c *gin.Context) {

}

type treatiseDeleteInput struct {
	Title string `json:"title"  binding:"required"`
}

// @Summary	admin DeleteTreatise
// @Security AdminAuth
// @Tags content
// @Description DeleteTreatise
// @ModuleID deleteTreatise
// @Accept  json
// @Produce  json
// @Param input body treatiseDeleteInput true "treatise info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/content/{id} [delete]
func (h *Handler) deleteTreatise(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		logger.Info(idParam)

		return
	}

	var inp treatiseDeleteInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	roles, ex := c.Get(roleCtx)
	if !ex {
		newResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err := h.services.Content.Delete(idParam, roles)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type createPartInput struct {
	ID          string `json:"id"  binding:"required"`
	TargetID    string
	Name        string `json:"name" binding:"required"`
	FullName    string `json:"full_name" binding:"required"`
	Description string `json:"description"  binding:"required"`
}

// @Summary	admin createPart
// @Security AdminAuth
// @Tags parts
// @Description createPart
// @ModuleID createPart
// @Accept  json
// @Produce  json
// @Param input body createPartInput true "part info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/content/{id}/parts [post]
func (h *Handler) createPart(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		logger.Info(idParam)

		return
	}
	var inp createPartInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	roles, ex := c.Get(roleCtx)
	if !ex {
		newResponse(c, http.StatusUnauthorized, "Unauthorized")
		return
	}

	err := h.services.Part.Create(inp.ID, idParam, inp.Name, inp.FullName, inp.Description, roles)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}
func (h *Handler) createProposition(c *gin.Context) {
}
