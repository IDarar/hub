package v1

import (
	"net/http"

	"github.com/IDarar/hub/internal/service"
	"github.com/IDarar/hub/pkg/logger"
	"github.com/gin-gonic/gin"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.POST("/sign-up", h.userSignUp)
		users.POST("/sign-in", h.userSignIn)
		users.POST("/auth/refresh", h.userRefresh)

		useractions := users.Group("/", h.userIdentity)

		{
			userContent := useractions.Group("/content")
			{
				userContent.POST("", h.addUserTreatise)
				userContent.PUT("/:id", h.updateUserTreatise)
				userContent.POST("/rate", h.rateTreatise)
				userContent.DELETE("/rate", h.deleteRateTreatise)

			}
			userParts := useractions.Group("/parts")
			{
				userParts.POST("", h.addUserPart)
				userParts.PUT("/:id", h.updateUserPart)
				userParts.POST("/rate", h.ratePart)
				userParts.DELETE("/rate", h.deelteRatePart)

			}
			userPropositions := useractions.Group("/propositions")
			{
				userPropositions.POST("", h.addUserProposition)
				userPropositions.PUT("/:id", h.updateUserProposition)
				userPropositions.POST("/rate", h.rateProposition)
				userPropositions.DELETE("/rate", h.deleteRateProposition)

			}

		}

	}
}

type signUpInput struct {
	Name     string `json:"name,omitempty"  binding:"required,min=2,max=64"`
	Email    string `json:"email,omitempty"  binding:"required,email,max=64"`
	Password string `json:"password,omitempty" binding:"required,min=4,max=64"`
}
type signInInput struct {
	Name     string `json:"name,omitempty"  binding:"required,min=2,max=64"`
	Password string `json:"password,omitempty" binding:"required,min=4,max=64"`
}

// @Summary User SignUp
// @Tags user-auth
// @Description create user account
// @ModuleID userSignUp
// @Accept  json
// @Produce  json
// @Param input body signUpInput true "sign up info"
// @Success 201 {string} string "ok"
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-up [post]
func (h *Handler) userSignUp(c *gin.Context) {
	var inp signUpInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.services.User.SignUp(c.Request.Context(), service.SignUpInput{
		Name:     inp.Name,
		Email:    inp.Email,
		Password: inp.Password,
	}); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type tokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

// @Summary user SignIn
// @Tags user-auth
// @Description user sign in
// @ModuleID user
// @Accept  json
// @Produce  json
// @Param input body signInInput true "sign up info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/sign-in [post]
func (h *Handler) userSignIn(c *gin.Context) {
	var inp signInInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.User.SignIn(c.Request.Context(), service.SignInInput{
		Name:     inp.Name,
		Password: inp.Password,
	})
	if err != nil {
		if err == service.ErrUserNotFound {
			newResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

type refreshInput struct {
	Token string `json:"token" binding:"required"`
}

// @Summary User Refresh Tokens
// @Tags user-auth
// @Description users refresh tokens
// @Accept  json
// @Produce  json
// @Param input body refreshInput true "token info"
// @Success 200 {object} tokenResponse
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/auth/refresh [post]
func (h *Handler) userRefresh(c *gin.Context) {
	var inp refreshInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	res, err := h.services.User.RefreshTokens(inp.Token)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, tokenResponse{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	})
}

type addTreatiseInput struct {
	TargetTreatise string `json:"target_treatise,omitempty" binding:"required"`
}

// @Summary	user addUserTreatise
// @Security UsersAuth
// @Tags UserContent
// @Description addUserTreatise
// @ModuleID user
// @Accept  json
// @Produce  json
// @Param input body addTreatiseInput true "content info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/content [post]
func (h *Handler) addUserTreatise(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	logger.Info(userID)

	var inp addTreatiseInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	logger.Info("USERID ", userID)

	err := h.services.User.AddTreatise(service.AddTreatiseInput{
		TargetTreatise: inp.TargetTreatise,
	}, userID)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type updateUserTreatise struct {
	TargetTreatise string `json:"target_treatise,omitempty" binding:"required"`
	Status         string `json:"status,omitempty"`
	IsCompleted    *bool  `json:"is_completed,omitempty"`
}

// @Summary	user updateUserTreatise
// @Security UsersAuth
// @Tags UserContent
// @Description updateUserTreatise
// @ModuleID user
// @Accept  json
// @Produce  json
// @Param input body updateUserTreatise true "content info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/content/{id} [put]
func (h *Handler) updateUserTreatise(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	idParam := c.Param("id")

	if idParam == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		logger.Info(idParam)

		return
	}

	var inp updateUserTreatise

	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if inp.IsCompleted == nil {
		logger.Info("is nil")
	}
	if inp.Status == "" && inp.IsCompleted == nil {
		newResponse(c, http.StatusBadRequest, "nil values, nothing to update")
		return
	}

	err := h.services.User.UpdateTreatise(service.UpdateUserTreatise{TargetTreatise: inp.TargetTreatise,
		Status:      inp.Status,
		IsCompleted: inp.IsCompleted},
		userID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type addPropositionInput struct {
	TargetProposition string `json:"target_proposition,omitempty" binding:"required"`
}

// @Summary	user addUserProposition
// @Security UsersAuth
// @Tags UserContent
// @Description addUserProposition
// @ModuleID user
// @Accept  json
// @Produce  json
// @Param input body addPropositionInput true "prop info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/propositions/ [post]
func (h *Handler) addUserProposition(c *gin.Context) {
	userID, _ := c.Get(userCtx)

	var inp addPropositionInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.User.AddProposition(service.AddPropositionInput{
		TargetProposition: inp.TargetProposition,
	}, userID)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type updateUserProposition struct {
	TargetProposition string `json:"target_proposition,omitempty" binding:"required"`
	Status            string `json:"status,omitempty"`
	IsCompleted       *bool  `json:"is_completed,omitempty"`
}

// @Summary	user updateUserProposition
// @Security UsersAuth
// @Tags UserContent
// @Description updateUserProposition
// @ModuleID user
// @Accept  json
// @Produce  json
// @Param input body updateUserProposition true "prop info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/propositions/{id} [put]
func (h *Handler) updateUserProposition(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	idParam := c.Param("id")

	if idParam == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		logger.Info(idParam)

		return
	}

	var inp updateUserProposition

	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	if inp.IsCompleted == nil {
		logger.Info("is nil")
	}
	if inp.Status == "" && inp.IsCompleted == nil {
		newResponse(c, http.StatusBadRequest, "nil values, nothing to update")
		return
	}

	err := h.services.User.UpdateProposition(service.UpdateUserProposition{TargetProposition: inp.TargetProposition,
		Status:      inp.Status,
		IsCompleted: inp.IsCompleted},
		userID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type addPartInput struct {
	TargetPart string `json:"target_part,omitempty" binding:"required"`
}

// @Summary	user addUserPart
// @Security UsersAuth
// @Tags UserContent
// @Description addUserPart
// @ModuleID user
// @Accept  json
// @Produce  json
// @Param input body addPartInput true "part info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/parts/ [post]
func (h *Handler) addUserPart(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	logger.Info(userID)

	var inp addPartInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	logger.Info("USERID ", userID)

	err := h.services.User.AddPart(service.AddPartInput{
		TargetPart: inp.TargetPart,
	}, userID)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type updateUserPart struct {
	TargetPart  string `json:"target_part,omitempty" binding:"required"`
	Status      string `json:"status,omitempty"`
	IsCompleted *bool  `json:"is_completed,omitempty"`
}

// @Summary	user updateUserPart
// @Security UsersAuth
// @Tags UserContent
// @Description updateUserPart
// @ModuleID user
// @Accept  json
// @Produce  json
// @Param input body updateUserPart true "part info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/parts/{id} [put]
func (h *Handler) updateUserPart(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	idParam := c.Param("id")

	if idParam == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		logger.Info(idParam)

		return
	}

	var inp updateUserPart

	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if inp.IsCompleted == nil {
		logger.Info("is nil")
	}
	if inp.Status == "" && inp.IsCompleted == nil {
		newResponse(c, http.StatusBadRequest, "nil values, nothing to update")
		return
	}

	err := h.services.User.UpdatePart(service.UpdateUserPart{TargetPart: inp.TargetPart,
		Status:      inp.Status,
		IsCompleted: inp.IsCompleted},
		userID)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

type rateInput struct {
	Target string `json:"target,omitempty" binding:"required"`
	Type   string `json:"type,omitempty" binding:"required"`
	Value  int    `json:"value,omitempty" binding:"required"`
}

// @Summary	user rateTreatise
// @Security UsersAuth
// @Tags Rates
// @Description rateTreatise
// @ModuleID Rates
// @Accept  json
// @Produce  json
// @Param input body rateInput true "rate info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/content/rate [post]
func (h *Handler) rateTreatise(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	logger.Info(userID)

	var inp rateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	logger.Info("USERID ", userID)

	err := h.services.User.RateTreatise(service.RateInput{
		Target: inp.Target,
		Value:  inp.Value,
		Type:   inp.Type,
	}, userID)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary	user ratePart
// @Security UsersAuth
// @Tags Rates
// @Description ratePart
// @ModuleID Rates
// @Accept  json
// @Produce  json
// @Param input body rateInput true "rate info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/parts/rate [post]
func (h *Handler) ratePart(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	logger.Info(userID)

	var inp rateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	logger.Info("USERID ", userID)

	err := h.services.User.RatePart(service.RateInput{
		Target: inp.Target,
		Value:  inp.Value,
		Type:   inp.Type,
	}, userID)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary	user rateProposition
// @Security UsersAuth
// @Tags Rates
// @Description rateProposition
// @ModuleID Rates
// @Accept  json
// @Produce  json
// @Param input body rateInput true "rate info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/propositions/rate [post]
func (h *Handler) rateProposition(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	logger.Info(userID)

	var inp rateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	logger.Info("USERID ", userID)

	err := h.services.User.RateProposition(service.RateInput{
		Target: inp.Target,
		Value:  inp.Value,
		Type:   inp.Type,
	}, userID)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary	user deleteRateTreatise
// @Security UsersAuth
// @Tags Rates
// @Description deleteRateTreatise
// @ModuleID Rates
// @Accept  json
// @Produce  json
// @Param input body rateInput true "rate info"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /users/content/rate [delete]
func (h *Handler) deleteRateTreatise(c *gin.Context) {
	userID, _ := c.Get(userCtx)
	logger.Info(userID)

	var inp rateInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	logger.Info("USERID ", userID)

	err := h.services.User.DeleteRateTreatise(service.RateInput{
		Target: inp.Target,
		Value:  inp.Value,
		Type:   inp.Type,
	}, userID)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusCreated)
}
func (h *Handler) deelteRatePart(c *gin.Context) {

}
func (h *Handler) deleteRateProposition(c *gin.Context) {

}
