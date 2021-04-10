package service

import (
	"strings"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/logger"
)

type PropositionsService struct {
	repo repository.Propositions
}

func NewPropositionsService(repo repository.Propositions) *PropositionsService {
	return &PropositionsService{
		repo: repo,
	}

}
func (s *PropositionsService) Create(id, TargetID, name, description, explanation, text string, roles interface{}) error {
	if err := checkRigths(roles, "admin"); err != nil {
		logger.Error(err)
		return err
	}

	proposition := domain.Proposition{ID: id, TargetID: strings.ToUpper(TargetID), Name: name, Explanation: explanation, Text: text, Description: description}
	if err := s.repo.Create(proposition); err != nil {
		logger.Error(err)
		return err
	}
	return nil

}
func (s *PropositionsService) Delete(id string, roles interface{}) error {
	return nil
}
