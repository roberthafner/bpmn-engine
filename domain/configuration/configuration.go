package configuration

import (
	"github.com/roberthafner/bpmn-engine/domain/model/command"
	"github.com/roberthafner/bpmn-engine/domain/model/database"
	"github.com/roberthafner/bpmn-engine/domain/model/manager"
	"github.com/roberthafner/bpmn-engine/domain/service"
)

type EngineConfiguration struct {
	db database.Database
	commandExecutor command.CommandExecutor
	deploymentEntityManager manager.DeploymentEntityManager
	DeploymentService service.DeploymentService
}

func NewEngineConfiguration() EngineConfiguration {

	db := database.NewDatabase()

	deploymentEntityManager := manager.NewDeploymentEntityManager(db)

	commandContext := command.CommandContext{
		deploymentEntityManager,
	}

	commandExecutor := command.NewDefaultCommandExecutor(commandContext)

	return EngineConfiguration{
		db: db,
		commandExecutor: commandExecutor,
		deploymentEntityManager: deploymentEntityManager,
		DeploymentService: service.NewDeploymentService(commandExecutor),
	}
}
