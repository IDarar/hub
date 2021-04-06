package http

import (
	"hub/internal/service"
	v1 "hub/internal/transport/http/v1"

	"github.com/gin-gonic/gin"
)

//1.06 21.13(another)
//Struct Hanlder takes all interfaces of service
type Handler struct {
	usersService  service.Users
	adminsService service.Admins
}

func NewHandler(usersService service.Users, adminsService service.Admins) *Handler {
	return &Handler{
		usersService:  usersService,
		adminsService: adminsService,
	}
}
func (h *Handler) Init() *gin.Engine {

	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")

	})
	h.initAPI(router)
	return router
}
func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.usersService, h.adminsService)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
