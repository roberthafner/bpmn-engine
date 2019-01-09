package command

import (
	"github.com/roberthafner/bpmn-engine/domain/model"
	"github.com/roberthafner/bpmn-engine/domain/model/manager"
	"time"
)

type Command interface {
	Execute(ctx CommandContext) interface{}
}

type CommandContext struct {
	DeploymentEntityManager manager.DeploymentEntityManager
}

type CommandExecutor interface {
	Execute(c Command) interface{}
}

type DefaultCommandExecutor struct {
	commandContext CommandContext
}

func (dce DefaultCommandExecutor) Execute(cmd Command) interface{} {
	return cmd.Execute(dce.commandContext)
}

func NewDefaultCommandExecutor(cc CommandContext) DefaultCommandExecutor {
	return DefaultCommandExecutor{commandContext: cc}
}

type DeployCmd struct {
	Deployment model.DeploymentEntity
}

func NewDeployCmd(d model.DeploymentEntity) DeployCmd {
	return DeployCmd {Deployment: d}
}

func (dc DeployCmd) Execute(ctx CommandContext) interface{} {
	d := dc.Deployment
	d.DeploymentTime = time.Now()
	dem := ctx.DeploymentEntityManager
	dem.Insert(d)
	return d
}
