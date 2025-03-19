package internal

import (
	"actions-back/config"
	actionsController "actions-back/internal/controller/actions"
	authController "actions-back/internal/controller/auth"
	actionsRepository "actions-back/internal/repository/actions"
	authRepository "actions-back/internal/repository/auth"
	actionsService "actions-back/internal/services/actions"
	authService "actions-back/internal/services/auth"

	"github.com/uptrace/bun"
)

type Container struct {
	DB                *bun.DB
	ActionsController *actionsController.ActionsController
	AuthController    *authController.AuthController
}

func NewContainer() *Container {
	database := config.NewDatabaseConfig()
	actionsRepository := actionsRepository.NewActionsRepository(database)
	actionsService := actionsService.NewActionsService(actionsRepository)
	actionsController := actionsController.NewActionsController(actionsService)

	authRepository := authRepository.NewUserRepository(database)
	authService := authService.NewAuthService(authRepository)
	authController := authController.NewAuthController(authService)

	return &Container{
		DB:                database,
		ActionsController: actionsController,
		AuthController:    authController,
	}
}
