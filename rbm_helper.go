package rbm

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

type RBMHelper struct {
	agentId string
	token   *oauth2.Token
	debug   bool
}

func GetRBMHelperInstanceFromFile(agentId string, filename string) (*RBMHelper, error) {
	token, err := GetAgentTokenFromFile(filename)
	if err != nil {
		return nil, fmt.Errorf("fail to get token: %v", err)
	}
	return &RBMHelper{
		agentId: agentId,
		token:   token,
		debug:   false,
	}, nil
}

func (h *RBMHelper) SetDebug(d bool) {
	h.debug = d
}

func (h *RBMHelper) SendEvent(msisdn string, event string, messageId string) error {

	bearer := "Bearer " + h.token.AccessToken

	eventId := uuid.New().String() + "a"
	url := "https://rcsbusinessmessaging.googleapis.com/v1/phones/" + msisdn + "/agentEvents?eventId=" + eventId

	var param = fmt.Sprintf(`{
		"eventType": "%s",
		"messageId": "%s"
	}`, event, messageId)
	var payload = bytes.NewBufferString(param)

	r, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", bearer)

	client := http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	debug := os.Getenv("DEBUG") == "1"

	if debug && res.StatusCode != http.StatusOK {
		PrintBodyRequest(res)
	}

	return err
}

func (h *RBMHelper) SendMessage(msisdn string, message []byte) (string, int, error) {
	bearer := "Bearer " + h.token.AccessToken

	messageId := uuid.New().String() + "a"
	url := fmt.Sprintf("https://rcsbusinessmessaging.googleapis.com/v1/phones/%s/agentMessages?messageId=%s&agentId=%s", msisdn, messageId, h.agentId)

	bodyReader := bytes.NewReader(message)

	r, err := http.NewRequest("POST", url, bodyReader)
	if err != nil {
		return messageId, 500, err
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", bearer)

	client := http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return messageId, 500, err
	}

	defer res.Body.Close()

	// debug := os.Getenv("DEBUG") == "1"

	// bodyBytes, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error get body", err)
	// }
	// fmt.Println(string(bodyBytes))

	if h.debug &&
		res.StatusCode != http.StatusOK &&
		res.StatusCode != http.StatusNotFound &&
		res.StatusCode != http.StatusForbidden {
		PrintBodyRequest(res)
	}

	return messageId, res.StatusCode, nil
}

func (h *RBMHelper) SendMultipleMessage(msisdn string, messages [][]byte) (string, int, error) {
	if len(messages) == 0 {
		return "", -1, fmt.Errorf("message empty")
	}

	var statusList []int
	var messageId string

	for _, message := range messages {
		rbmId, status, err := SendMessage(h.agentId, h.token.AccessToken, msisdn, message)
		messageId = rbmId
		if err != nil {
			return messageId, 500, err
		}
		statusList = append(statusList, status)
		if status != 200 {
			break
		}
	}
	return messageId, statusList[len(statusList)-1], nil
}

func (h *RBMHelper) CapabilityCheck(msisdn string) (int, error) {

	bearer := "Bearer " + h.token.AccessToken

	requestId := uuid.New().String() + "a"

	url := fmt.Sprintf("https://rcsbusinessmessaging.googleapis.com/v1/phones/%s/capabilities?requestId=%s&agentId=%s", msisdn, requestId, h.agentId)

	r, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 500, err
	}
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", bearer)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return 500, err
	}

	defer res.Body.Close()

	debug := os.Getenv("DEBUG") == "1"

	if debug && res.StatusCode != http.StatusOK {
		PrintBodyRequest(res)
	}

	return res.StatusCode, nil
}
