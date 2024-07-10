package rbm

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

func GetAgentFromFile(filename string) (map[string]string, error) {
	plan, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("fail to read file: %v", err)
	}

	var data map[string]string

	err = json.Unmarshal(plan, &data)
	if err != nil {
		return nil, fmt.Errorf("fail to parse json: %v", err)
	}

	return data, nil
}

func GetAgentConfig(agent map[string]string) *jwt.Config {
	config := &jwt.Config{
		Email:      agent["client_email"],
		PrivateKey: []byte(agent["private_key"]),
		Scopes: []string{
			"https://www.googleapis.com/auth/rcsbusinessmessaging",
		},
		TokenURL: google.JWTTokenURL,
	}
	return config
}

func GetAgentTokenFromFile(filename string) (*oauth2.Token, error) {
	agent, err := GetAgentFromFile(filename)
	if err != nil {
		return nil, err
	}

	config := GetAgentConfig(agent)

	token, err := config.TokenSource(context.TODO()).Token()
	if err != nil {
		return nil, err
	}

	return token, nil
}
