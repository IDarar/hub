package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/pkg/auth"
	"github.com/IDarar/hub/pkg/hash"
	"github.com/IDarar/hub/pkg/logger"
)

//39.58
type UserService struct {
	repo         repository.Users
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager

	accessTokenTTL         time.Duration
	refreshTokenTTL        time.Duration
	verificationCodeLength int
}

func NewUsersService(repo repository.Users, hasher hash.PasswordHasher,
	tokenManager auth.TokenManager, accessTTL, refreshTTL time.Duration,
	verificationCodeLength int) *UserService {
	return &UserService{
		repo:                   repo,
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

	return s.createSession(user.ID)
}
func (s *UserService) CreateMark(domain.UserProposition, [3]interface{}) error {

	return nil
}
func (s *UserService) createSession(userId int) (Tokens, error) {
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

	err = s.repo.SetSession(userId, session)
	return res, err
}
func (s *UserService) GetRoleById(Userid int) ([]string, error) {
	roles, err := s.repo.GetRoleByID(Userid)
	if err != nil {
		return roles, errors.New("dont have enough rights")
	}
	return roles, nil
}
