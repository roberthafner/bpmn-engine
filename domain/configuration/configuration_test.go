package configuration

import (
	"github.com/roberthafner/bpmn-engine/domain/model"
	"os"
	"testing"
)

func TestConfiguration(t *testing.T) {
	definition :=
		`<definitions>
			<process id="1" name="test">
				<startEvent id="2" name="start"/>
				<sequenceFlow id="3" name="sequence flow 3" sourceRef="2" targetRef="4"/>
				<userTask id="4" name="usertask"/>"
				<sequenceFlow id="5" name="sequence flow 5" sourceRef="4" targetRef="6"/>
				<endEvent id="6" name="end"/>
			</process>
		</definitions>`


	engineConfiguration := NewEngineConfiguration()
	d := model.NewDeployment("123", "TestDeployment")
	dr := model.NewDeploymentResource("1234", "123", "resource1", []byte(definition))
	d.Resources = append(d.Resources, dr)


	deploymentService := engineConfiguration.DeploymentService
	deploymentService.CreateDeployment(d)
	path, _ := os.Getwd()
	os.Remove(path + "/burrow.db")
}