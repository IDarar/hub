package service

import (
	"strings"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/logger"
)

type PartsService struct {
	repo  repository.Parts
	users User
}

func NewPartsService(repo repository.Parts, users User) *PartsService {
	return &PartsService{
		repo:  repo,
		users: users,
	}

}
func (s *PartsService) Create(id, TargetID, name, fullname, description string, roles interface{}) error {
	roles, err := s.users.GetRoleById(roles.(int))

	if err != nil {
		logger.Error(err)
		return err
	}

	if err := checkRigths(roles, "admin"); err != nil {
		logger.Error(err)
		return err
	}

	part := domain.Part{ID: id, TargetID: strings.ToUpper(TargetID), Name: name, FullName: fullname, Description: description}
	if err := s.repo.Create(part); err != nil {
		logger.Error(err)
		return err
	}
	return nil

}
func (s *PartsService) Update(inp PartUpdateInput, roles interface{}) error {
	roles, err := s.users.GetRoleById(roles.(int))
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := checkRigths(roles, "admin"); err != nil {
		logger.Error(err)
		return err
	}

	part := domain.Part{
		ID:          inp.ID,
		TargetID:    strings.ToUpper(inp.TargetID),
		Name:        inp.Name,
		FullName:    inp.FullName,
		Description: inp.Description,
	}

	if err := s.repo.Update(part, inp.CreateLiterature, inp.DeleteLiterature); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (s *PartsService) Delete(id string, roles interface{}) error {
	return nil
}
