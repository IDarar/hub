package http

import (
	_ "github.com/IDarar/hub/docs"
	"github.com/IDarar/hub/internal/elasticsearch"
	"github.com/IDarar/hub/internal/service"
	v1 "github.com/IDarar/hub/internal/transport/http/v1"
	"github.com/IDarar/hub/pkg/auth"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//Struct Hanlder takes all interfaces of service
type Handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
	indexer      *elasticsearch.Indexers
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager, indexer *elasticsearch.Indexers) *Handler {
	return &Handler{
		services:     services,
		tokenManager: tokenManager,
		indexer:      indexer,
	}
}

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")

	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.initAPI(router)
	return router
}
func (h *Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.services, h.tokenManager, h.indexer)
	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}
