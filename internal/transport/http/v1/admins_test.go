package v1

import (
	"bytes"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/IDarar/hub/internal/service"
	"github.com/IDarar/hub/internal/service/mocks"
	"github.com/IDarar/hub/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestHandler_createProposition(t *testing.T) {
	logger.Info("Start")
	type mockBehavior func(s *mocks.MockPropositions, input service.CreateProposition, roles interface{})
	tests := []struct {
		name                string
		inputBody           string
		inputProp           service.CreateProposition
		idParam             string
		mockBehavior        mockBehavior
		roles               interface{}
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			inputBody: `{"id":"123","name":"TYR","description":"123","explanation":"test-prop","text":"132"}`,
			idParam:   "RLP",
			inputProp: service.CreateProposition{
				ID:          "123",
				Name:        "TYR",
				TargetID:    "RLP",
				Description: "123",
				Explanation: "test-prop",
				Text:        "132",
			},
			roles: []string{"admin"},
			mockBehavior: func(s *mocks.MockPropositions, input service.CreateProposition, roles interface{}) {
				s.EXPECT().Create(input, roles)
			},
			expectedStatusCode: 401,
		},
		{
			name:      "OK",
			inputBody: `{"id":"123","name":"TYR","description":"123","explanation":"test-prop","text":"132"}`,
			idParam:   "RLP",
			inputProp: service.CreateProposition{
				ID:          "123",
				Name:        "TYR",
				TargetID:    "RLP",
				Description: "123",
				Explanation: "test-prop",
				Text:        "132",
			},
			roles: []string{"admin"},
			mockBehavior: func(s *mocks.MockPropositions, input service.CreateProposition, roles interface{}) {
				s.EXPECT().Create(input, roles)
			},
			expectedStatusCode: 500,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			prop := mocks.NewMockPropositions(c)
			testCase.mockBehavior(prop, testCase.inputProp, testCase.roles)

			services := &service.Services{Propositions: prop}
			handler := Handler{services: services}

			r := gin.New()
			r.POST("/create-prop/:id", handler.createProposition)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", fmt.Sprintf("/create-prop/%s", testCase.idParam),
				bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}

}
