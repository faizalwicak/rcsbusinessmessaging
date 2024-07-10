package rbm

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/google/uuid"
)

func SendEvent(token string, msisdn string, event string, messageId string) error {

	bearer := "Bearer " + token

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

func SendMessage(agentId string, token string, msisdn string, message []byte) (string, int, error) {
	bearer := "Bearer " + token

	messageId := uuid.New().String() + "a"
	url := fmt.Sprintf("https://rcsbusinessmessaging.googleapis.com/v1/phones/%s/agentMessages?messageId=%s&agentId=%s", msisdn, messageId, agentId)

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

	debug := os.Getenv("DEBUG") == "1"

	// bodyBytes, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("error get body", err)
	// }
	// fmt.Println(string(bodyBytes))

	if debug &&
		res.StatusCode != http.StatusOK &&
		res.StatusCode != http.StatusNotFound &&
		res.StatusCode != http.StatusForbidden {
		PrintBodyRequest(res)
	}

	return messageId, res.StatusCode, nil
}

func SendMultipleMessage(agentId string, token string, msisdn string, messages [][]byte) (string, int, error) {
	if len(messages) == 0 {
		return "", -1, fmt.Errorf("message empty")
	}

	var statusList []int
	var messageId string

	for _, message := range messages {
		rbmId, status, err := SendMessage(agentId, token, msisdn, message)
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

func CapabilityCheck(token string, msisdn string) (int, error) {

	bearer := "Bearer " + token

	requestId := uuid.New().String() + "a"

	url := "https://rcsbusinessmessaging.googleapis.com/v1/phones/" + msisdn + "/capabilities?requestId=" + requestId

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
