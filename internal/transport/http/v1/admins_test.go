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
func TestHandler_updateTreatise(t *testing.T) {
	logger.Info("Start")
	type mockBehavior func(s *mocks.MockContent, input service.TreatiseUpdateInput, userID string)
	tests := []struct {
		name                string
		userID              string
		inputBody           string
		inputTreat          service.TreatiseUpdateInput
		idParam             string
		mockBehavior        mockBehavior
		expectedStatusCode  int
		expectedRequestBody string
	}{
		{
			name:      "OK",
			userID:    "1",
			inputBody: `{"date": "1656","description": "NOTLASTFORVER","title": "On imp"}`,
			idParam:   "RLP",
			inputTreat: service.TreatiseUpdateInput{
				ID:          "RLP",
				Title:       "On imp",
				Date:        "1656",
				Description: "NOTLASTFORVER",
			},
			mockBehavior: func(s *mocks.MockContent, input service.TreatiseUpdateInput, userID string) {
				s.EXPECT().Update(input, userID).Return(nil)
			},
			expectedStatusCode: 201,
		},
		{
			name:      "invalid data",
			userID:    "1",
			idParam:   "TP",
			inputBody: `{"descrip": "NOTLASTFORVER"}`,
			inputTreat: service.TreatiseUpdateInput{
				ID: "TP",
			},
			mockBehavior: func(s *mocks.MockContent, input service.TreatiseUpdateInput, userID string) {
				s.EXPECT().Update(input, userID)
			},
			expectedStatusCode: 400,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			treatise := mocks.NewMockContent(c)

			testCase.mockBehavior(treatise, testCase.inputTreat, testCase.userID)

			services := &service.Services{Content: treatise}
			handler := Handler{services: services}

			r := gin.New()

			r.PUT("/content/:id", func(c *gin.Context) {
				c.Set(userCtx, testCase.userID)

			}, handler.updateTreatise)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("PUT", fmt.Sprintf("/content/%s", testCase.idParam),
				bytes.NewBufferString(testCase.inputBody))
			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedStatusCode, w.Code)

		})
	}

}
