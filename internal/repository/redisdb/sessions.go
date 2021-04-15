package redisdb

import (
	"context"
	"time"

	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/internal/domain"
	"github.com/IDarar/hub/pkg/logger"
	"github.com/go-redis/redis/v8"
)

type SessionsRepo struct {
	rdb *redis.Client
	cfg *config.Config
	ctx context.Context
}

func NewSessionRepo(rdb *redis.Client, cfg *config.Config) *SessionsRepo {
	return &SessionsRepo{
		rdb: rdb,
		cfg: cfg,
	}
}

func (r *SessionsRepo) SetSession(userId int, session domain.Session) error {
	ctx := context.Background()
	start := time.Now()
	end := time.Now().Add(r.cfg.Auth.JWT.RefreshTokenTTL)

	difference := end.Sub(start)
	err := r.rdb.Set(ctx, session.RefreshToken, userId, difference)
	if err.Err() != nil {
		logger.Error(err.Err())
		return err.Err()
	}

	return nil
}
func (r *StudentsRepo) GetByRefreshToken(ctx context.Context, schoolId primitive.ObjectID, refreshToken string) (domain.Student, error) {
	var student domain.Student
	if err := r.db.FindOne(ctx, bson.M{"session.refreshToken": refreshToken, "schoolId": schoolId,
		"session.expiresAt": bson.M{"$gt": time.Now()}}).Decode(&student); err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Student{}, ErrUserNotFound
		}

		return domain.Student{}, err
	}

	return student, nil
}
