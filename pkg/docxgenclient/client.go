package docxgenclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type DocxgenClient struct {
	Host string
	Port int
}

func New(host string, port int) *DocxgenClient {
	return &DocxgenClient{Host: host, Port: port}
}

func (dc *DocxgenClient) GetReport(req Request) ([]byte, error) {
	json, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("cant marshall docx-gen req to json: %w", err)
	}

	request, err := http.NewRequest("POST", "http://"+dc.Host+":"+strconv.Itoa(dc.Port)+"/gen",
		bytes.NewBuffer(json))
	if err != nil {
		return nil, fmt.Errorf("cant create report request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("cant get docx-gen js service response: %w", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cant read chart docx-gen service response body: %w", err)
	}

	return body, nil
}

func (dc *DocxgenClient) GetShortReport(req RequestShortReport) ([]byte, error) {
	json, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("cant marshall docx-gen req to json: %w", err)
	}

	request, err := http.NewRequest("POST", "http://"+dc.Host+":"+strconv.Itoa(dc.Port)+"/genShort",
		bytes.NewBuffer(json))
	if err != nil {
		return nil, fmt.Errorf("cant create short report request: %v", err)
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("cant get docx-gen js service response: %w", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("cant read chart docx-gen service response body: %w", err)
	}
	return body, nil
}
