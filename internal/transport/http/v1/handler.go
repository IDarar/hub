package v1

import (
	"hub/internal/service"

	"github.com/gin-gonic/gin"
)

//51.13 31.43
type Handler struct {
	usersService  service.Users
	adminsService service.Admins
}

func NewHandler(usersService service.Users, adminsService service.Admins) *Handler {
	return &Handler{usersService: usersService, adminsService: adminsService}
}
func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initUsersRoutes(v1)
		h.initAdminsRoutes(v1)
	}
}
func newResponse(c *gin.Context, code int, message interface{}) {
	c.String(code, "text", message)

}
