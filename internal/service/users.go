package service

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/auth"
	"github.com/IDarar/hub/pkg/hash"
	"github.com/IDarar/hub/pkg/logger"
)

type UserService struct {
	repo         repository.Users
	sessions     repository.Sessions
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL         time.Duration
	refreshTokenTTL        time.Duration
	verificationCodeLength int
}

func NewUsersService(repo repository.Users, sessions repository.Sessions, hasher hash.PasswordHasher,
	tokenManager auth.TokenManager, accessTTL, refreshTTL time.Duration,
	verificationCodeLength int) *UserService {
	return &UserService{
		repo:                   repo,
		sessions:               sessions,
		hasher:                 hasher,
		verificationCodeLength: verificationCodeLength,
		tokenManager:           tokenManager,
		accessTokenTTL:         accessTTL,
		refreshTokenTTL:        refreshTTL,
	}
}

func (s *UserService) SignUp(ctx context.Context, input SignUpInput) error {
	user := domain.User{
		Name:         input.Name,
		Password:     s.hasher.Hash(input.Password),
		Email:        input.Email,
		RegisteredAt: time.Now(),
		LastVisitAt:  time.Now(),
		Session:      domain.Session{RefreshToken: "", ExpiresAt: time.Now().Add(s.refreshTokenTTL)},
	}
	logger.Info(user)
	if err := s.repo.Create(ctx, user); err != nil {
		return err
	}
	return nil

}
func (s *UserService) SignIn(ctx context.Context, input SignInInput) (Tokens, error) {
	user, err := s.repo.GetByCredentials(ctx, input.Name, s.hasher.Hash(input.Password))
	if err != nil {
		if err == repository.ErrUserNotFound {
			return Tokens{}, ErrUserNotFound
		}
		return Tokens{}, err
	}

	return s.createSession(user.ID, "")
}
func (s *UserService) RefreshTokens(refreshToken string) (Tokens, error) {
	//check if token exists and then pass it to delete and set new
	uID, err := s.sessions.GetIDByRefreshToken(refreshToken)
	if err != nil {
		return Tokens{}, err
	}

	return s.createSession(uID, refreshToken)
}

func (s *UserService) GetRoleById(Userid int) ([]string, error) {
	roles, err := s.repo.GetRoleByID(Userid)
	if err != nil {
		return roles, errors.New("dont have enough rights")
	}
	return roles, nil
}

func (s *UserService) CreateMark(domain.UserProposition, [3]interface{}) error {

	return nil
}
func (s *UserService) createSession(userId int, revoketoken string) (Tokens, error) {
	var (
		res Tokens
		err error
	)

	res.AccessToken, err = s.tokenManager.NewJWT(fmt.Sprint(userId), s.accessTokenTTL)
	if err != nil {
		return res, err
	}

	res.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return res, err
	}

	session := domain.Session{
		RefreshToken: res.RefreshToken,
		ExpiresAt:    time.Now().Add(s.refreshTokenTTL),
	}

	err = s.sessions.SetSession(userId, session, revoketoken)
	return res, err
}
func (s *UserService) AddTreatise(inp AddTreatiseInput, userID interface{}) error {
	logger.Info("userID ", userID)

	treatise := domain.UserTreatise{
		TargetTreatise: strings.ToUpper(inp.TargetTreatise),
		UserID:         userID.(int)}
	logger.Info(treatise)

	if err := s.repo.AddTreatise(treatise); err != nil {
		logger.Error(err)
		return err
	}
	return nil

}
func (s *UserService) UpdateTreatise(inp UpdateUserTreatise, userID interface{}) error {
	if inp.IsCompleted == nil {
		logger.Info("is nil")
	}
	treatise := domain.UserTreatise{
		TargetTreatise: strings.ToUpper(inp.TargetTreatise),
		Status:         inp.Status,
		UserID:         userID.(int),
		IsCompleted:    inp.IsCompleted}
	if err := s.repo.UpdateTreatise(treatise); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (s *UserService) AddProposition(inp AddPropositionInput, userID interface{}) error {
	logger.Info("userID ", userID)

	prop := domain.UserProposition{
		TargetProposition: strings.ToUpper(inp.TargetProposition),
		UserID:            userID.(int),
		Status:            "In progress"}
	logger.Info(prop)

	if err := s.repo.AddProposition(prop); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (s *UserService) UpdateProposition(inp UpdateUserProposition, userID interface{}) error {
	if inp.IsCompleted == nil {
		logger.Info("is nil")
	}
	prop := domain.UserProposition{TargetProposition: strings.ToUpper(inp.TargetProposition),
		Status:      inp.Status,
		UserID:      userID.(int),
		IsCompleted: inp.IsCompleted}
	if err := s.repo.UpdateProposition(prop); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (s *UserService) AddPart(inp AddPartInput, userID interface{}) error {
	logger.Info("userID ", userID)

	part := domain.UserPart{TargetPart: strings.ToUpper(inp.TargetPart), UserID: userID.(int)}
	logger.Info(part)

	if err := s.repo.AddPart(part); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (s *UserService) UpdatePart(inp UpdateUserPart, userID interface{}) error {
	if inp.IsCompleted == nil {
		logger.Info("is nil")
	}
	part := domain.UserPart{TargetPart: strings.ToUpper(inp.TargetPart),
		Status:      inp.Status,
		UserID:      userID.(int),
		IsCompleted: inp.IsCompleted}
	if err := s.repo.UpdatePart(part); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (s *UserService) RateTreatise(rateinp RateInput, userID interface{}) error {
	logger.Info("userID ", userID)

	treatise, err := checkContentRateType(rateinp)
	if err != nil {
		logger.Error(err)
		return err
	}
	treatise.UserID = userID.(int)
	treatise.TargetTreatise = rateinp.Target

	rate := domain.Rate{
		TargetID: treatise.TargetTreatise,
		UserID:   treatise.UserID,
		Value:    rateinp.Value,
		Type:     rateinp.Type,
	}
	logger.Info(treatise)

	if err := s.repo.RateTreatise(treatise, rate); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (s *UserService) RatePart(rateinp RateInput, userID interface{}) error {
	logger.Info("userID ", userID)

	part, err := checkPartRateType(rateinp)
	if err != nil {
		logger.Error(err)
		return err
	}
	part.UserID = userID.(int)
	part.TargetPart = rateinp.Target

	rate := domain.Rate{
		TargetID: part.TargetPart,
		UserID:   part.UserID,
		Value:    rateinp.Value,
		Type:     rateinp.Type,
	}
	logger.Info(part)

	if err := s.repo.RatePart(part, rate); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (s *UserService) RateProposition(rateinp RateInput, userID interface{}) error {
	logger.Info("userID ", userID)

	pr, err := checkPropositionRateType(rateinp)
	if err != nil {
		logger.Error(err)
		return err
	}
	pr.UserID = userID.(int)
	pr.TargetProposition = rateinp.Target

	rate := domain.Rate{
		TargetID: pr.TargetProposition,
		UserID:   pr.UserID,
		Value:    rateinp.Value,
		Type:     rateinp.Type,
	}
	logger.Info(pr)

	if err := s.repo.RateProposition(pr, rate); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func checkContentRateType(rate RateInput) (domain.UserTreatise, error) {
	treatise := domain.UserTreatise{}

	switch rate.Type {
	case "difficulty":
		treatise.DifficultyRate = rate.Value
		return treatise, nil
	case "importance":
		treatise.ImportanceRate = rate.Value
		return treatise, nil
	case "inconsistency":
		treatise.InconsistencyRate = rate.Value
		return treatise, nil
	default:
		return treatise, errors.New("invalid rate type")
	}
}
func checkPartRateType(rate RateInput) (domain.UserPart, error) {
	part := domain.UserPart{}

	switch rate.Type {
	case "difficulty":
		part.DifficultyRate = rate.Value
		return part, nil
	case "importance":
		part.ImportanceRate = rate.Value
		return part, nil
	case "inconsistency":
		part.InconsistencyRate = rate.Value
		return part, nil
	default:
		return part, errors.New("invalid rate type")
	}
}
func checkPropositionRateType(rate RateInput) (domain.UserProposition, error) {
	part := domain.UserProposition{}

	switch rate.Type {
	case "difficulty":
		part.DifficultyRate = rate.Value
		return part, nil
	case "importance":
		part.ImportanceRate = rate.Value
		return part, nil
	case "inconsistency":
		part.InconsistencyRate = rate.Value
		return part, nil
	default:
		return part, errors.New("invalid rate type")
	}
}
func checkContentRateTypeDelete(rate RateInput) (domain.UserTreatise, error) {
	tr := domain.UserTreatise{}

	switch rate.Type {
	case "difficulty":
		tr.DifficultyRate = rate.Value
		return tr, nil
	case "importance":
		tr.ImportanceRate = rate.Value
		return tr, nil
	case "inconsistency":
		tr.InconsistencyRate = rate.Value
		return tr, nil
	default:
		return tr, errors.New("invalid rate type")
	}
}
func checkPartRateTypeDelete(rate RateInput) (domain.UserPart, error) {
	part := domain.UserPart{}

	switch rate.Type {
	case "difficulty":
		part.DifficultyRate = rate.Value
		return part, nil
	case "importance":
		part.ImportanceRate = rate.Value
		return part, nil
	case "inconsistency":
		part.InconsistencyRate = rate.Value
		return part, nil
	default:
		return part, errors.New("invalid rate type")
	}
}
func checkPropositionRateTypeDelete(rate RateInput) (domain.UserProposition, error) {
	pr := domain.UserProposition{}

	switch rate.Type {
	case "difficulty":
		pr.DifficultyRate = rate.Value
		return pr, nil
	case "importance":
		pr.ImportanceRate = rate.Value
		return pr, nil
	case "inconsistency":
		pr.InconsistencyRate = rate.Value
		return pr, nil
	default:
		return pr, errors.New("invalid rate type")
	}
}
func (s *UserService) DeleteRateTreatise(rateinp RateInput, userID interface{}) error {
	logger.Info("userID ", userID)

	tr, err := checkContentRateTypeDelete(rateinp)
	if err != nil {
		logger.Error(err)
		return err
	}
	tr.UserID = userID.(int)
	tr.TargetTreatise = rateinp.Target

	rate := domain.Rate{
		TargetID: tr.TargetTreatise,
		UserID:   tr.UserID,
		Value:    rateinp.Value,
		Type:     rateinp.Type,
	}
	logger.Info(tr)

	if err := s.repo.DeleteRateTreatise(tr, rate); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (s *UserService) DeleteRatePart(rateinp RateInput, userID interface{}) error {
	logger.Info("userID ", userID)

	part, err := checkPartRateTypeDelete(rateinp)
	if err != nil {
		logger.Error(err)
		return err
	}
	part.UserID = userID.(int)
	part.TargetPart = rateinp.Target

	rate := domain.Rate{
		TargetID: part.TargetPart,
		UserID:   part.UserID,
		Value:    rateinp.Value,
		Type:     rateinp.Type,
	}
	logger.Info(part)

	if err := s.repo.DeleteRatePart(part, rate); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}

func (s *UserService) DeleteRateProposition(rateinp RateInput, userID interface{}) error {
	logger.Info("userID ", userID)

	pr, err := checkPropositionRateTypeDelete(rateinp)
	if err != nil {
		logger.Error(err)
		return err
	}
	pr.UserID = userID.(int)
	pr.TargetProposition = rateinp.Target

	rate := domain.Rate{
		TargetID: pr.TargetProposition,
		UserID:   pr.UserID,
		Value:    rateinp.Value,
		Type:     rateinp.Type,
	}
	logger.Info(pr)

	if err := s.repo.DeleteRateProposition(pr, rate); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
