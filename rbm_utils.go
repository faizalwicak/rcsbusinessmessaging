package rbm

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func GetRequestBody(res *http.Response) (string, error) {
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(bodyBytes), nil
}

func PrintBodyRequest(res *http.Response) {
	var log = logrus.New()
	log.Out = os.Stdout

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}

	bodyString, err := GetRequestBody(res)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(bodyString)
}

func StructToJson(data any) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return []byte{}
	}
	return b
}

func JsonToStruct(data []byte) map[string]any {
	var b map[string]any
	if err := json.Unmarshal(data, &b); err != nil {
		log.Println(err)
		return map[string]any{}
	}
	return b
}
