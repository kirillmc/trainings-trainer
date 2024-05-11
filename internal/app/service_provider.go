package app

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/kirillmc/platform_common/pkg/closer"
	"github.com/kirillmc/platform_common/pkg/db"
	"github.com/kirillmc/platform_common/pkg/db/pg"
	descAccess "github.com/kirillmc/trainings-auth/pkg/access_v1"
	"github.com/kirillmc/trainings-server/internal/api/moderator"
	"github.com/kirillmc/trainings-server/internal/api/training"
	"github.com/kirillmc/trainings-server/internal/client/rpc"
	"github.com/kirillmc/trainings-server/internal/client/rpc/access"
	"github.com/kirillmc/trainings-server/internal/config"
	"github.com/kirillmc/trainings-server/internal/config/env"
	"github.com/kirillmc/trainings-server/internal/interceptor"
	"github.com/kirillmc/trainings-server/internal/repository"
	moderRepo "github.com/kirillmc/trainings-server/internal/repository/moder"
	trainingRepo "github.com/kirillmc/trainings-server/internal/repository/training"
	"github.com/kirillmc/trainings-server/internal/service"
	moderService "github.com/kirillmc/trainings-server/internal/service/moder"
	trainingService "github.com/kirillmc/trainings-server/internal/service/training"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	accessConfig  config.AccessConfig

	dbClient          db.Client
	accessClient      rpc.AccessClient
	interceptorClient *interceptor.Interceptor

	trainingRepository     repository.TrainingRepository
	trainingService        service.TrainingService
	trainingImplementation *training.Implementation

	moderRepository     repository.ModerRepository
	moderService        service.ModerService
	moderImplementation *moderator.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		pgConfig, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgConfig = pgConfig
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		grpcConfig, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		s.grpcConfig = grpcConfig
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		httpConfig, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %v", err)
		}

		s.httpConfig = httpConfig
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		swaggerConfig, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %v", err)
		}

		s.swaggerConfig = swaggerConfig
	}

	return s.swaggerConfig
}

func (s *serviceProvider) AccessConfig() config.AccessConfig {
	if s.accessConfig == nil {
		accessConfig, err := env.NewAccessConfig()
		if err != nil {
			log.Fatalf("failed to get access configs: %v", err)
		}

		s.accessConfig = accessConfig
	}

	return s.accessConfig
}

func (s *serviceProvider) AccessClient() rpc.AccessClient {
	if s.accessClient == nil {
		cfg := s.AccessConfig()

		conn, err := grpc.Dial(
			cfg.Address(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			log.Fatalf("failed to connect to access: %v", err)
		}

		s.accessClient = access.NewAccessClient(descAccess.NewAccessV1Client(conn))
	}

	return s.accessClient
}

func (s *serviceProvider) InterceptorClient() *interceptor.Interceptor {
	if s.interceptorClient == nil {
		s.interceptorClient = &interceptor.Interceptor{
			Client: s.AccessClient(),
		}
	}

	return s.interceptorClient
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}

		closer.Add(cl.Close)
		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TrainingRepository(ctx context.Context) repository.TrainingRepository {
	if s.trainingRepository == nil {
		s.trainingRepository = trainingRepo.NewRepository(s.DBClient(ctx))
	}

	return s.trainingRepository
}

func (s *serviceProvider) ModerRepository(ctx context.Context) repository.ModerRepository {
	if s.moderRepository == nil {
		s.moderRepository = moderRepo.NewRepository(s.DBClient(ctx))
	}

	return s.moderRepository
}

func (s *serviceProvider) TrainingService(ctx context.Context) service.TrainingService {
	if s.trainingService == nil {
		s.trainingService = trainingService.NewService(s.TrainingRepository(ctx))
	}

	return s.trainingService
}

func (s *serviceProvider) ModerService(ctx context.Context) service.ModerService {
	if s.moderService == nil {
		s.moderService = moderService.NewService(s.ModerRepository(ctx))
	}

	return s.moderService
}

func (s *serviceProvider) TrainingImplementation(ctx context.Context) *training.Implementation {
	if s.trainingImplementation == nil {
		s.trainingImplementation = training.NewImplementation(s.TrainingService(ctx))
	}

	return s.trainingImplementation
}

func (s *serviceProvider) ModerImplementation(ctx context.Context) *moderator.Implementation {
	if s.moderImplementation == nil {
		s.moderImplementation = moderator.NewImplementation(s.ModerService(ctx))
	}

	return s.moderImplementation
}
