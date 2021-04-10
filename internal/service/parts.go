package service

import (
	"strings"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/logger"
)

type PartsService struct {
	repo repository.Parts
}

func NewPartsService(repo repository.Parts) *PartsService {
	return &PartsService{
		repo: repo,
	}

}
func (s *PartsService) Create(id, TargetID, name, fullname, description string, roles interface{}) error {
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
func (s *PartsService) Delete(id string, roles interface{}) error {
	return nil
}
