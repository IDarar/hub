// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/services.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	domain "github.com/IDarar/hub/internal/domain"
	service "github.com/IDarar/hub/internal/service"
	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// AddPart mocks base method.
func (m *MockUser) AddPart(inp service.AddPartInput, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPart", inp, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPart indicates an expected call of AddPart.
func (mr *MockUserMockRecorder) AddPart(inp, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPart", reflect.TypeOf((*MockUser)(nil).AddPart), inp, userID)
}

// AddProposition mocks base method.
func (m *MockUser) AddProposition(inp service.AddPropositionInput, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProposition", inp, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProposition indicates an expected call of AddProposition.
func (mr *MockUserMockRecorder) AddProposition(inp, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProposition", reflect.TypeOf((*MockUser)(nil).AddProposition), inp, userID)
}

// AddTreatise mocks base method.
func (m *MockUser) AddTreatise(inp service.AddTreatiseInput, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTreatise", inp, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTreatise indicates an expected call of AddTreatise.
func (mr *MockUserMockRecorder) AddTreatise(inp, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTreatise", reflect.TypeOf((*MockUser)(nil).AddTreatise), inp, userID)
}

// CreateMark mocks base method.
func (m *MockUser) CreateMark(arg0 domain.UserProposition, arg1 [3]interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMark", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateMark indicates an expected call of CreateMark.
func (mr *MockUserMockRecorder) CreateMark(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMark", reflect.TypeOf((*MockUser)(nil).CreateMark), arg0, arg1)
}

// DeleteRatePart mocks base method.
func (m *MockUser) DeleteRatePart(rate service.RateInput, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRatePart", rate, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRatePart indicates an expected call of DeleteRatePart.
func (mr *MockUserMockRecorder) DeleteRatePart(rate, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRatePart", reflect.TypeOf((*MockUser)(nil).DeleteRatePart), rate, userID)
}

// DeleteRateProposition mocks base method.
func (m *MockUser) DeleteRateProposition(rate service.RateInput, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRateProposition", rate, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRateProposition indicates an expected call of DeleteRateProposition.
func (mr *MockUserMockRecorder) DeleteRateProposition(rate, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRateProposition", reflect.TypeOf((*MockUser)(nil).DeleteRateProposition), rate, userID)
}

// DeleteRateTreatise mocks base method.
func (m *MockUser) DeleteRateTreatise(rate service.RateInput, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRateTreatise", rate, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRateTreatise indicates an expected call of DeleteRateTreatise.
func (mr *MockUserMockRecorder) DeleteRateTreatise(rate, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRateTreatise", reflect.TypeOf((*MockUser)(nil).DeleteRateTreatise), rate, userID)
}

// GetRoleById mocks base method.
func (m *MockUser) GetRoleById(id int) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRoleById", id)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRoleById indicates an expected call of GetRoleById.
func (mr *MockUserMockRecorder) GetRoleById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRoleById", reflect.TypeOf((*MockUser)(nil).GetRoleById), id)
}

// RatePart mocks base method.
func (m *MockUser) RatePart(rate service.RateInput, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RatePart", rate, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RatePart indicates an expected call of RatePart.
func (mr *MockUserMockRecorder) RatePart(rate, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RatePart", reflect.TypeOf((*MockUser)(nil).RatePart), rate, userID)
}

// RateProposition mocks base method.
func (m *MockUser) RateProposition(rate service.RateInput, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RateProposition", rate, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RateProposition indicates an expected call of RateProposition.
func (mr *MockUserMockRecorder) RateProposition(rate, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RateProposition", reflect.TypeOf((*MockUser)(nil).RateProposition), rate, userID)
}

// RateTreatise mocks base method.
func (m *MockUser) RateTreatise(rate service.RateInput, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RateTreatise", rate, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RateTreatise indicates an expected call of RateTreatise.
func (mr *MockUserMockRecorder) RateTreatise(rate, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RateTreatise", reflect.TypeOf((*MockUser)(nil).RateTreatise), rate, userID)
}

// RefreshTokens mocks base method.
func (m *MockUser) RefreshTokens(refreshToken string) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RefreshTokens", refreshToken)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RefreshTokens indicates an expected call of RefreshTokens.
func (mr *MockUserMockRecorder) RefreshTokens(refreshToken interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RefreshTokens", reflect.TypeOf((*MockUser)(nil).RefreshTokens), refreshToken)
}

// SignIn mocks base method.
func (m *MockUser) SignIn(ctx context.Context, input service.SignInInput) (service.Tokens, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", ctx, input)
	ret0, _ := ret[0].(service.Tokens)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserMockRecorder) SignIn(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUser)(nil).SignIn), ctx, input)
}

// SignUp mocks base method.
func (m *MockUser) SignUp(ctx context.Context, input service.SignUpInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", ctx, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserMockRecorder) SignUp(ctx, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUser)(nil).SignUp), ctx, input)
}

// UpdatePart mocks base method.
func (m *MockUser) UpdatePart(inp service.UpdateUserPart, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePart", inp, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePart indicates an expected call of UpdatePart.
func (mr *MockUserMockRecorder) UpdatePart(inp, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePart", reflect.TypeOf((*MockUser)(nil).UpdatePart), inp, userID)
}

// UpdateProposition mocks base method.
func (m *MockUser) UpdateProposition(inp service.UpdateUserProposition, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProposition", inp, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProposition indicates an expected call of UpdateProposition.
func (mr *MockUserMockRecorder) UpdateProposition(inp, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProposition", reflect.TypeOf((*MockUser)(nil).UpdateProposition), inp, userID)
}

// UpdateTreatise mocks base method.
func (m *MockUser) UpdateTreatise(inp service.UpdateUserTreatise, userID interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTreatise", inp, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTreatise indicates an expected call of UpdateTreatise.
func (mr *MockUserMockRecorder) UpdateTreatise(inp, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTreatise", reflect.TypeOf((*MockUser)(nil).UpdateTreatise), inp, userID)
}

// MockAdmin is a mock of Admin interface.
type MockAdmin struct {
	ctrl     *gomock.Controller
	recorder *MockAdminMockRecorder
}

// MockAdminMockRecorder is the mock recorder for MockAdmin.
type MockAdminMockRecorder struct {
	mock *MockAdmin
}

// NewMockAdmin creates a new mock instance.
func NewMockAdmin(ctrl *gomock.Controller) *MockAdmin {
	mock := &MockAdmin{ctrl: ctrl}
	mock.recorder = &MockAdminMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdmin) EXPECT() *MockAdminMockRecorder {
	return m.recorder
}

// GrantRole mocks base method.
func (m *MockAdmin) GrantRole(name, role string, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GrantRole", name, role, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// GrantRole indicates an expected call of GrantRole.
func (mr *MockAdminMockRecorder) GrantRole(name, role, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GrantRole", reflect.TypeOf((*MockAdmin)(nil).GrantRole), name, role, roles)
}

// RevokeRole mocks base method.
func (m *MockAdmin) RevokeRole(user *domain.User, role string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RevokeRole", user, role)
	ret0, _ := ret[0].(error)
	return ret0
}

// RevokeRole indicates an expected call of RevokeRole.
func (mr *MockAdminMockRecorder) RevokeRole(user, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeRole", reflect.TypeOf((*MockAdmin)(nil).RevokeRole), user, role)
}

// MockContent is a mock of Content interface.
type MockContent struct {
	ctrl     *gomock.Controller
	recorder *MockContentMockRecorder
}

// MockContentMockRecorder is the mock recorder for MockContent.
type MockContentMockRecorder struct {
	mock *MockContent
}

// NewMockContent creates a new mock instance.
func NewMockContent(ctrl *gomock.Controller) *MockContent {
	mock := &MockContent{ctrl: ctrl}
	mock.recorder = &MockContentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContent) EXPECT() *MockContentMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockContent) Create(id, title, date, description string, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", id, title, date, description, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockContentMockRecorder) Create(id, title, date, description, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockContent)(nil).Create), id, title, date, description, roles)
}

// Delete mocks base method.
func (m *MockContent) Delete(id string, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockContentMockRecorder) Delete(id, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockContent)(nil).Delete), id, roles)
}

// Update mocks base method.
func (m *MockContent) Update(inp service.TreatiseUpdateInput, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", inp, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockContentMockRecorder) Update(inp, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockContent)(nil).Update), inp, roles)
}

// MockPart is a mock of Part interface.
type MockPart struct {
	ctrl     *gomock.Controller
	recorder *MockPartMockRecorder
}

// MockPartMockRecorder is the mock recorder for MockPart.
type MockPartMockRecorder struct {
	mock *MockPart
}

// NewMockPart creates a new mock instance.
func NewMockPart(ctrl *gomock.Controller) *MockPart {
	mock := &MockPart{ctrl: ctrl}
	mock.recorder = &MockPartMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPart) EXPECT() *MockPartMockRecorder {
	return m.recorder
}

// AddToFavourite mocks base method.
func (m *MockPart) AddToFavourite(fav domain.Favourite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToFavourite", fav)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToFavourite indicates an expected call of AddToFavourite.
func (mr *MockPartMockRecorder) AddToFavourite(fav interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToFavourite", reflect.TypeOf((*MockPart)(nil).AddToFavourite), fav)
}

// Create mocks base method.
func (m *MockPart) Create(id, TargetID, name, fullname, description string, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", id, TargetID, name, fullname, description, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPartMockRecorder) Create(id, TargetID, name, fullname, description, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPart)(nil).Create), id, TargetID, name, fullname, description, roles)
}

// Delete mocks base method.
func (m *MockPart) Delete(id string, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPartMockRecorder) Delete(id, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPart)(nil).Delete), id, roles)
}

// RemoveFromFavourite mocks base method.
func (m *MockPart) RemoveFromFavourite(fav domain.Favourite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromFavourite", fav)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromFavourite indicates an expected call of RemoveFromFavourite.
func (mr *MockPartMockRecorder) RemoveFromFavourite(fav interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromFavourite", reflect.TypeOf((*MockPart)(nil).RemoveFromFavourite), fav)
}

// Update mocks base method.
func (m *MockPart) Update(inp service.PartUpdateInput, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", inp, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPartMockRecorder) Update(inp, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPart)(nil).Update), inp, roles)
}

// MockPropositions is a mock of Propositions interface.
type MockPropositions struct {
	ctrl     *gomock.Controller
	recorder *MockPropositionsMockRecorder
}

// MockPropositionsMockRecorder is the mock recorder for MockPropositions.
type MockPropositionsMockRecorder struct {
	mock *MockPropositions
}

// NewMockPropositions creates a new mock instance.
func NewMockPropositions(ctrl *gomock.Controller) *MockPropositions {
	mock := &MockPropositions{ctrl: ctrl}
	mock.recorder = &MockPropositionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPropositions) EXPECT() *MockPropositionsMockRecorder {
	return m.recorder
}

// AddToFavourite mocks base method.
func (m *MockPropositions) AddToFavourite(fav domain.Favourite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToFavourite", fav)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToFavourite indicates an expected call of AddToFavourite.
func (mr *MockPropositionsMockRecorder) AddToFavourite(fav interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToFavourite", reflect.TypeOf((*MockPropositions)(nil).AddToFavourite), fav)
}

// Create mocks base method.
func (m *MockPropositions) Create(prop service.CreateProposition, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", prop, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockPropositionsMockRecorder) Create(prop, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockPropositions)(nil).Create), prop, roles)
}

// Delete mocks base method.
func (m *MockPropositions) Delete(id string, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockPropositionsMockRecorder) Delete(id, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPropositions)(nil).Delete), id, roles)
}

// RemoveFromFavourite mocks base method.
func (m *MockPropositions) RemoveFromFavourite(fav domain.Favourite) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveFromFavourite", fav)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveFromFavourite indicates an expected call of RemoveFromFavourite.
func (mr *MockPropositionsMockRecorder) RemoveFromFavourite(fav interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveFromFavourite", reflect.TypeOf((*MockPropositions)(nil).RemoveFromFavourite), fav)
}

// Update mocks base method.
func (m *MockPropositions) Update(inp service.UpdatePropositionInput, roles interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", inp, roles)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockPropositionsMockRecorder) Update(inp, roles interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockPropositions)(nil).Update), inp, roles)
}
