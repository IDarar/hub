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
				//TODO change handlers
				userContent.PUT("/:id", h.updateTreatise)
				userContent.DELETE("/:id", h.deleteTreatise)
				userContent.POST("/:id/parts", h.createPart)
				userContent.POST("/:id/proposition", h.createProposition) //For treatises without parts division

			}
			userParts := useractions.Group("/parts")
			{
				userParts.POST("/:id")
				userParts.PUT("/:id", h.updatePart)
				userParts.DELETE("/:id", h.deletePart)
				userParts.POST("/:id/proposition", h.createProposition) //For treatises divided into parts
			}
			userPropositions := useractions.Group("/propositions")
			{
				userPropositions.PUT("/:id", h.updateProposition)
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
	UserID int `json:"user_id" binding:"required"`

	TargetTreatise string `json:"target_treatise" binding:"required"`
}

func (h *Handler) addUserTreatise(c *gin.Context) {
	idParam := c.Param("id")
	userID, _ := c.Get(userCtx)
	logger.Info(userID)
	if idParam == "" {
		newResponse(c, http.StatusBadRequest, "empty id param")
		logger.Info(idParam)

		return
	}

	var inp addTreatiseInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.User.AddTreatise(service.AddTreatiseInput{
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

type updateTreatise struct {
	UserID int

	TargetTreatise string
	Status         string

	DifficultyRate    int
	ImportanceRate    int
	InconsistencyRate int

	Progress    int
	IsCompleted bool
}
