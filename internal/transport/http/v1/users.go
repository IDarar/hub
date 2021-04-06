package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("/users") //51.23
	{
		users.POST("/sign-up", h.userSignUp)

		users.POST("/sign-in", h.userSignIn)

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

//25.45
func (h *Handler) userSignUp(c *gin.Context) {
	var inp signInInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	newResponse(c, 200, "works")
}

func (h *Handler) userSignIn(c *gin.Context) {

	var inp signInInput
	if err := c.BindJSON(&inp); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	//further logik
}
