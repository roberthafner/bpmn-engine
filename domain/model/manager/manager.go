package manager

import (
	"github.com/roberthafner/bpmn-engine/domain/model"
	"github.com/roberthafner/bpmn-engine/domain/model/database"
)

type EntityManager interface {
	Insert(e model.Entity)
	Update(e model.Entity)
	Delete(e model.Entity)
}


type DeploymentEntityManager struct {
	db database.Database
}

func NewDeploymentEntityManager(db database.Database) DeploymentEntityManager {
	dem := DeploymentEntityManager{}
	dem.db = db
	return dem
}

func (dem DeploymentEntityManager) Insert(e model.Entity) {
	d := e.(model.DeploymentEntity)
	dem.db.Insert(d)

	resources := d.Resources
	for _, r := range resources {
		r.DeploymentId = d.Id
		dem.db.Insert(r)
	}
}

func (dem DeploymentEntityManager) Update(e model.Entity) {
	d := e.(model.DeploymentEntity)
	dem.db.Update(d)
}

func (dem DeploymentEntityManager) Delete(e model.Entity) {
	d := e.(model.DeploymentEntity)
	dem.db.Delete(d)
}
