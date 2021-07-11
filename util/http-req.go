package util

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

type HttpOption struct {
	Url     string
	Method  string
	Header  []string    // key, value, key, value...
	Payload interface{} // struct or string
}

func HttpReq(option *HttpOption) ([]byte, error) {
	url := option.Url
	method := option.Method
	payload := strings.NewReader("")

	switch option.Payload.(type) {
		case string:
			payload = strings.NewReader(option.Payload.(string))
		case nil:
		default:
			payloadBytes, err := json.Marshal(option.Payload)
			if err != nil {
				return nil, err
			}
			payload = strings.NewReader(string(payloadBytes))
	}

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return nil, err
	}

	headerLen := len(option.Header)
	if headerLen%2 != 0 {
		headerLen--
	}
	for i := 0; i < headerLen; i += 2 {
		req.Header.Add(option.Header[i], option.Header[i + 1])
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	return body, err
}
