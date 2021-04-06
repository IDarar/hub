package service

import "github.com/IDarar/hub/internal/domain"

//all interfaces there are described
type UserActions interface {
	CreateMark(domain.UserProposition, [3]interface{}) error
}
type Services struct {
	User UserActions
}

//TODO 39.47
func NewServices() *Services {
	return &Services{}
}
