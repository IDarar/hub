package service

import (
	"strings"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/logger"
)

type ContentService struct {
	repo repository.Content
}

func NewContentService(repo repository.Content) *ContentService {
	return &ContentService{
		repo: repo,
	}

}
func (s *ContentService) Create(id, title, date, description string, roles interface{}) error {
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
func (s *ContentService) Delete(id, title string, roles interface{}) error {
	if err := checkRigths(roles, "admin"); err != nil {
		logger.Error(err)
		return err
	}
	treatise := domain.Treatise{ID: strings.ToUpper(id), Title: title}
	if err := s.repo.Delete(treatise); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
