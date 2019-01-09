package service

import (
	"github.com/roberthafner/bpmn-engine/domain/model"
	"github.com/roberthafner/bpmn-engine/domain/model/command"
)

type DeploymentService interface {
	CreateDeployment(d model.DeploymentEntity) model.DeploymentEntity
}

func NewDeploymentService(ce command.CommandExecutor) DeploymentService {
	return deploymentService {
		commandExecutor: ce,
	}
}

type deploymentService struct {
	commandExecutor command.CommandExecutor
}

func (ds deploymentService) CreateDeployment(d model.DeploymentEntity) model.DeploymentEntity {
	cmd := command.NewDeployCmd(d)
	return ds.commandExecutor.Execute(cmd).(model.DeploymentEntity)
}

