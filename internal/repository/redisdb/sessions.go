package redisdb

import (
	"context"
	"strconv"
	"time"

	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"
	"github.com/go-redis/redis/v8"
)

type SessionsRepo struct {
	rdb *redis.Client
	cfg *config.Config
}

func NewSessionRepo(rdb *redis.Client, cfg *config.Config) *SessionsRepo {
	return &SessionsRepo{
		rdb: rdb,
		cfg: cfg,
	}
}

func (r *SessionsRepo) SetSession(userId int, session domain.Session, revoketoken string) error {
	logger.Info("u id ", userId)
	ctx := context.Background()
	start := time.Now()
	end := time.Now().Add(r.cfg.Auth.JWT.RefreshTokenTTL)

	difference := end.Sub(start)
	err := r.rdb.Set(ctx, session.RefreshToken, userId, difference)
	if err.Err() != nil {
		logger.Error(err.Err())
		return err.Err()
	}

	if revoketoken != "" {
		r.rdb.Del(ctx, revoketoken)
	}

	logger.Info("new session ", session.RefreshToken)

	return nil
}
func (r *SessionsRepo) GetIDByRefreshToken(refreshToken string) (int, error) {
	logger.Info("refresh token to check ", refreshToken)

	ctx := context.Background()
	uIDstr, err := r.rdb.Get(ctx, refreshToken).Result()

	if err != nil {
		logger.Error(err)
		return 0, err
	}
	uID, _ := strconv.Atoi(uIDstr)
	logger.Info("u id ", uID)

	return int(uID), nil
}

//TODO in long future
//Add map of slices with user's sessions info
func (r *SessionsRepo) GetAllUserSessions(uID int) ([]string, error) {
	return nil, nil
}
