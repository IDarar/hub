package v1

import (
	"testing"

	"github.com/IDarar/hub/internal/service"
	"github.com/IDarar/hub/pkg/auth"
	"github.com/gin-gonic/gin"
)

func TestHandler_updateUserPart(t *testing.T) {
	type fields struct {
		services     *service.Services
		tokenManager auth.TokenManager
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &Handler{
				services:     tt.fields.services,
				tokenManager: tt.fields.tokenManager,
			}
			h.updateUserPart(tt.args.c)
		})
	}
}
