package repository

import "hub/internal/domain"

//all db interfaces there are described 44.20
//and struct repositories named with all interfaces in the end defined
type Users interface {
	GetUserByID(int) (*domain.User, error)
	CreateMark(map[int]int)
}
type Admins interface {
	//TODO
	grantRole()
	revokeRole()
}
type Content interface {
}

type Repositories struct {
	Users  Users
	Admins Admins
}
