package service

import (
	"strings"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/logger"
)

type ContentService struct {
	repo  repository.Content
	users User
}

func NewContentService(repo repository.Content, users User) *ContentService {
	return &ContentService{
		repo:  repo,
		users: users,
	}

}
func (s *ContentService) Create(id, title, date, description string, roles interface{}) error {
	roles, err := s.users.GetRoleById(roles.(int))
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := checkRigths(roles, "admin"); err != nil {
		logger.Error(err)
		return err
	}
	treatise := domain.Treatise{ID: id, Title: title, Date: date, Description: description}
	if err := s.repo.Create(treatise); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (s *ContentService) Update(inp TreatiseUpdateInput, roles interface{}) error {
	roles, err := s.users.GetRoleById(roles.(int))
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := checkRigths(roles, "admin"); err != nil {
		logger.Error(err)
		return err
	}
	treatise := domain.Treatise{ID: inp.ID, Title: inp.Title, Date: inp.Date, Description: inp.Description}
	if err := s.repo.Update(treatise); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (s *ContentService) Delete(id string, roles interface{}) error {
	roles, err := s.users.GetRoleById(roles.(int))
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := checkRigths(roles, "admin"); err != nil {
		logger.Error(err)
		return err
	}
	treatise := domain.Treatise{ID: strings.ToUpper(id)}
	if err := s.repo.Delete(treatise); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
