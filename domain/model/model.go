package model

import (
	"time"
)

type Entity interface {
	GetId() string
}

type DeploymentEntity struct {
	Id string `db:"DEPLOYMENT_ID"`
	Name string `db:"DEPLOYMENT_NAME"`
	DeploymentTime time.Time `db:"DEPLOYMENT_TIME"`
	Resources []DeploymentResource //`db:"-"`
}

func (de DeploymentEntity) GetId() string {
	return de.Id
}

func NewDeployment(id string, name string) DeploymentEntity {
	d := DeploymentEntity {Id:id, Name:name}
	return d
}

type DeploymentResource struct {
	Id string `db:"RESOURCE_ID"`
	DeploymentId string `db:"DEPLOYMENT_ID"`
	Name string `db:"RESOURCE_NAME"`
	Bytes []byte `db:"RESOURCE_BYTES"`
}

func (dr DeploymentResource) GetId() string {
	return dr.Id
}

func NewDeploymentResource(id string, deploymentId string, name string, bytes []byte) DeploymentResource {
	return DeploymentResource { Id: id, Name : name, DeploymentId: deploymentId, Bytes: bytes}
}
