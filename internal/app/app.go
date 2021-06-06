package app

import (
	"context"
	"flag"
	"os"

	"github.com/IDarar/hub/internal/config"
	"github.com/IDarar/hub/internal/elasticsearch"
	"github.com/IDarar/hub/internal/repository"
	"github.com/IDarar/hub/internal/repository/postgres"
	"github.com/IDarar/hub/internal/repository/redisdb"
	"github.com/IDarar/hub/internal/server"
	"github.com/IDarar/hub/internal/service"
	grpcv1 "github.com/IDarar/hub/internal/transport/grpc/v1"
	"github.com/IDarar/hub/internal/transport/http"
	"github.com/IDarar/hub/pkg/auth"
	"github.com/IDarar/hub/pkg/hash"
	"github.com/IDarar/hub/pkg/logger"
)

var ctx = context.Background()

// @title Hub
// @version 0.001
// @description Hub

// @host subjless.space/api/
// @BasePath /api/v1/

// @securityDefinitions.apikey AdminAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey UsersAuth
// @in header
// @name Authorization
func Run(configPath string) {
	envInit()

	cfg, err := config.Init(configPath)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("Config is: ", cfg)
	tokenManager, err := auth.NewManager(cfg.Auth.JWT.SigningKey)
	if err != nil {
		logger.Error(err)
		return
	}
	db, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		logger.Error(err)
		return
	}

	logger.Info("connected to postgres")
	rdb, err := redisdb.NewRedisDB(cfg)
	if err != nil {
		logger.Error(err)
		return
	}
	defer rdb.Close()

	rdb.Ping(ctx)
	logger.Info("connected to redis")

	hasher := hash.NewSHA1Hasher(cfg.Auth.PasswordSalt)

	repos := repository.NewRepositories(db, rdb, cfg)

	notCl := grpcv1.InitNotificationServiceClient(cfg)

	/*_, err = notCl.Client.NotificationCreate(ctx, &pb.Notification{Type: "forum", To: "124", From: "8643", Where: "ERT", Content: "15", Time: timestamppb.Now()})
	if err != nil {
		logger.Info(err)
		return
	}*/

	services := service.NewServices(service.Deps{
		Repos:                  repos,
		Hasher:                 hasher,
		AccessTokenTTL:         cfg.Auth.JWT.AccessTokenTTL,
		RefreshTokenTTL:        cfg.Auth.JWT.RefreshTokenTTL,
		TokenManager:           tokenManager,
		NotificationGrpcClient: *notCl,
	})
	elastic, err := elasticsearch.NewElasticSearch(*cfg)
	if err != nil {
		logger.Error(err)
		//return
	}
	logger.Info("connected to elasticsearch")

	indexers := elasticsearch.NewIndexer(elastic)
	handlers := http.NewHandler(services, tokenManager, indexers)

	srv := server.NewServer(cfg, handlers.Init())

	srv.Run()
}

//is used to check the deploy environment. if local, there will be set env variables
//that in other situations are passed to dockerfile in .env
func envInit() {
	e := flag.Bool("env", false, "run app local?")

	flag.Parse()
	logger.Info("deploy env is ", *e)

	if *e {
		os.Setenv("POSTGRES_PORT", "5432")
		os.Setenv("POSTGRES_HOST", "localhost")

		os.Setenv("POSTGRES_USER", "root")
		os.Setenv("POSTGRES_DBNAME", "root")

		os.Setenv("POSTGRES_PASSWORD", "secret")

		//TODO env for test db
		//os.Setenv("DATABASE_URL", "user=postgres dbname=hub_tests password=123 sslmode=disabled")

		os.Setenv("REDIS_ADDR", "localhost:6379")
		os.Setenv("REDIS_PASSWORD", "")

		os.Setenv("REDIS_DB", "0")

		os.Setenv("PASSWORD_SALT", "1234")

		os.Setenv("JWT_SIGNINGKEY", "signing_key")
	}
}

//docker pull 508084547447.dkr.ecr.us-west-2.amazonaws.com/hub:latest
//aws ecr get-login-password --region us-west-2 | docker login --username AWS --password-stdin 508084547447.dkr.ecr.region.amazonaws.com
//508084547447.dkr.ecr.us-east-1.amazonaws.com/hub:9154f112c026a245ee9c70d814f8de2e5f843e4b
//

//TODO, how to choose which service to connect?
func initExtService() bool {
	e := flag.Bool("ext", false, "connect to external services?")
	flag.Parse()
	if *e {
		return true
	} else {
		return false
	}
}
