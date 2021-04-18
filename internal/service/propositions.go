package service

import (
	"strings"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/logger"
)

type PropositionsService struct {
	repo  repository.Propositions
	users User
}

func NewPropositionsService(repo repository.Propositions, users User) *PropositionsService {
	return &PropositionsService{
		repo:  repo,
		users: users,
	}

}
func (s *PropositionsService) Create(prop CreateProposition, roles interface{}) error {
	roles, err := s.users.GetRoleById(roles.(int))
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := checkRigths(roles, "admin"); err != nil {
		logger.Error(err)
		return err
	}
	proposition := domain.Proposition{
		ID:          prop.ID,
		TargetID:    strings.ToUpper(prop.TargetID),
		Name:        prop.Name,
		Description: prop.Description,
		Explanation: prop.Explanation,
		Text:        prop.Text}
	if err := s.repo.Create(proposition); err != nil {
		logger.Error(err)
		return err
	}

	return nil

}
func (s *PropositionsService) AddToFavourite(fav domain.Favourite) error {
	logger.Info(fav)
	err := s.repo.AddToFavourite(fav)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (s *PropositionsService) RemoveFromFavourite(fav domain.Favourite) error {
	err := s.repo.AddToFavourite(fav)
	if err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (s *PropositionsService) Delete(id string, roles interface{}) error {
	return nil
}
func (s *PropositionsService) Update(inp UpdatePropositionInput, roles interface{}) error {
	roles, err := s.users.GetRoleById(roles.(int))
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := checkRigths(roles, "admin"); err != nil {
		logger.Error(err)
		return err
	}
	proposition := domain.Proposition{
		ID:          inp.ID,
		TargetID:    strings.ToUpper(inp.TargetID),
		Name:        inp.Name,
		Description: inp.Description,
		Explanation: inp.Explanation,
		Text:        inp.Text,
	}

	if err := s.repo.Update(proposition, inp.CreateReferences, inp.DeleteReferences, inp.CreateNotes, inp.DeleteNotes); err != nil {
		logger.Error(err)
		return err
	}
	/*if len(inp.CreateReferences) == 0 && len(inp.DeleteReferences) == 0 {
		logger.Info("no references to update, retunr nil err")
		return nil
	}*/

	return nil
}
