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

func TestGetAgentTokenFromFile(t *testing.T) {
	_, err := GetAgentTokenFromFile(filepath.Join("temp", "credential-test-2.json"))
	if err != nil {
		t.Errorf("fail to generate token: %v", err)
	}
}

func TestGenerateTextMessage(t *testing.T) {
	message := GetTextMessage("hello", []any{})
	assert := assert.New(t)
	assert.Equal(string(message), "{\"contentMessage\":{\"text\":\"hello\"}}")
}

func TestSendMessage(t *testing.T) {
	message := GetTextMessage("hello", []any{})

	token, err := GetAgentTokenFromFile(filepath.Join("temp", "credential-test-2.json"))
	if err != nil {
		t.Errorf("fail to generate token: %v", err)
	}

	_, status, err := SendMessage("", token.AccessToken, "", message)
	if err != nil {
		t.Errorf("fail to send message: %v", err)
	}

	if status != 200 {
		t.Errorf("fail to send message: message status: %d", status)
	}
}

func TestSendMessageWithInstance(t *testing.T) {
	message := GetTextMessage("hello from instance", []any{})

	rbmHelper, err := GetRBMHelperInstanceFromFile("telkomsel_khxk1xjf_agent", filepath.Join("temp", "credential-test-2.json"))
	if err != nil {
		t.Errorf("fail to get rbm instance: %v", err)
	}

	rbmHelper.SendMessage("+6287778680696", message)
}
