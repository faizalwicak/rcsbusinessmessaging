package rbm

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAgentFromFile(t *testing.T) {
	assert := assert.New(t)

	agent, err := GetAgentFromFile(filepath.Join("temp", "credential-test.json"))
	if err != nil {
		t.Errorf("fail to get file: %v", err)
	}

	assert.Contains(agent, "project_id")
	assert.Contains(agent, "client_email")
	assert.Contains(agent, "private_key")

	assert.Equal(agent["project_id"], "project_id")
	assert.Equal(agent["client_email"], "client_email")
	assert.Equal(agent["private_key"], "private_key")

	agentConfig := GetAgentConfig(agent)
	if err != nil {
		t.Errorf("fail to get file: %v", err)
	}

	assert.Equal(agentConfig.Email, agent["client_email"])
}
