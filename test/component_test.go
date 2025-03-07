package test

import (
	// "strings"
	"testing"

	helper "github.com/cloudposse/test-helpers/pkg/atmos/component-helper"
	// "github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
)

type ComponentSuite struct {
	helper.TestSuite
}

func (s *ComponentSuite) TestBasic() {
	const component = "example/basic"
	const stack = "default-test"
	const awsRegion = "us-east-2"

	defer s.DestroyAtmosComponent(s.T(), component, stack, nil)
	options, _ := s.DeployAtmosComponent(s.T(), component, stack, nil)
	assert.NotNil(s.T(), options)

	s.DriftTest(component, stack, nil)
}


func (s *ComponentSuite) TestEnabledFlag() {
	const component = "example/disabled"
	const stack = "default-test"
	s.VerifyEnabledFlag(component, stack, nil)
}

func TestRunSuite(t *testing.T) {
	suite := new(ComponentSuite)

	// TODO: Add dependency for the VPC component
	// suite.AddDependency(t, "vpc", "default-test", nil)

	// TODO: Add dependency for the DNS component
	// With unique inputs to avoid conflicts

	// subdomain := strings.ToLower(random.UniqueId())
	// inputs := map[string]interface{}{
	// 	"zone_config": []map[string]interface{}{
	// 		{
	// 			"subdomain": subdomain,
	// 			"zone_name": "components.cptest.test-automation.app",
	// 		},
	// 	},
	// }
	// suite.AddDependency(t, "dns-delegated", "default-test", &inputs)

	helper.Run(t, suite)
}
