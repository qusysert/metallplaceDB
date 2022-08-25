package converter

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Converter struct {
	Host string
	Port int
}

func New(host string, port int) *Converter {
	return &Converter{Host: host, Port: port}
}

func (c *Converter) GetPDFFromHTML(htmlBytes bytes.Buffer, path string) error {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "file.html")
	io.Copy(part, bytes.NewReader(htmlBytes.Bytes()))
	writer.Close()

	req, err := http.NewRequest("POST", "http://"+c.Host+":"+strconv.Itoa(c.Port)+"/convert?auth=key", body)
	if err != nil {
		return fmt.Errorf("cant create request: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("cant convert: %w", err)
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		str, _ := ioutil.ReadAll(rsp.Body)
		return fmt.Errorf("request failed with response code: %d, %s", rsp.StatusCode, str)
	}

	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("cant create pdf file: %w", err)
	}
	_, err = io.Copy(file, rsp.Body)
	return err
}