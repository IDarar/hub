package v1

import (
	"github.com/IDarar/hub/internal/elasticsearch"
	"github.com/IDarar/hub/internal/service"
	"github.com/IDarar/hub/pkg/auth"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1", h.RequestIndexer)
	{
		h.initUsersRoutes(v1)
		h.initAdminsRoutes(v1)
	}
}
