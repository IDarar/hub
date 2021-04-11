package v1

import (
	"net/http"

	"github.com/IDarar/hub/internal/service"
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
		parts := admins.Group("/parts")
		{
			/*parts.POST("/:id")
			parts.PUT("/:id")
			parts.PUT("/:id")*/
			//For treatises divided into parts
			parts.POST("/:id/proposition", h.createProposition)
		} /*
			propositions := content.Group("/propositions")
			{
				propositions.POST("/:id")
				propositions.PUT("/:id")
				propositions.PUT("/:id")
			}*/

	}

}

type RoleInput struct {
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
// @Param input body RoleInput true "role granting info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/roles/grant-role [post]
func (h *Handler) grantRole(c *gin.Context) {
	userID := c.Param(userCtx)

	var inp RoleInput

	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Admin.GrantRole(inp.UserName, inp.Role, userID)
	if err != nil {
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

//TODO

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
	userID, _ := c.Get(userCtx)

	var inp treatiseCreateInput

	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Content.Create(inp.ID, inp.Title, inp.Date, inp.Description, userID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type treatiseUpdateInput struct {
	ID          string
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

// @Summary	admin updateTreatise
// @Security AdminAuth
// @Tags content
// @Description updateTreatise
// @ModuleID updateTreatise
// @Accept  json
// @Produce  json
// @Param input body treatiseUpdateInput true "treatise update info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/content/{id} [put]
func (h *Handler) updateTreatise(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	idParam := c.Param("id")

	if idParam == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		logger.Info(idParam)

		return
	}
	var inp treatiseUpdateInput

	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Content.Update(service.TreatiseUpdateInput{
		ID:          idParam,
		Date:        inp.Date,
		Title:       inp.Title,
		Description: inp.Description,
	},
		userID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
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
	userID, _ := c.Get(userCtx)

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

	err := h.services.Content.Delete(idParam, userID)
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
	userID, _ := c.Get(userCtx)

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

	err := h.services.Part.Create(inp.ID, idParam, inp.Name, inp.FullName, inp.Description, userID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type createPropositionInput struct {
	ID          string `json:"id"  binding:"required"`
	TargetID    string
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"  binding:"required"`
	Explanation string `json:"explanation"  binding:"required"`
	Text        string `json:"text"  binding:"required"`
}

// @Summary	admin createProposition
// @Security createProposition
// @Tags propositions
// @Description createProposition
// @ModuleID createProposition
// @Accept  json
// @Produce  json
// @Param input body createPropositionInput true "proposition info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /admins/content/{id}/proposition/ [post]
func (h *Handler) createProposition(c *gin.Context) {
	idParam := c.Param("id")
	userID, _ := c.Get(userCtx)
	logger.Info(userID)
	if idParam == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		logger.Info(idParam)

		return
	}

	var inp createPropositionInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Propositions.Create(service.CreateProposition{
		ID:          inp.ID,
		TargetID:    idParam,
		Name:        inp.Name,
		Description: inp.Description,
		Explanation: inp.Explanation,
		Text:        inp.Text},
		userID)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}
func (h *Handler) updateProposition(c *gin.Context) {
}
